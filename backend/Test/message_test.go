package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/pkg/auth"
	"backend/pkg/chat"
	"backend/pkg/follow"
	"backend/pkg/handler"
	l "backend/pkg/log"
	"backend/pkg/messages"
	"backend/pkg/structs"

	uuid "github.com/satori/go.uuid"
)

func TestMessage(t *testing.T) {
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}

	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	follow.FollowUser(result1.UserId, result2.UserId, database)
	follow.FollowUser(result2.UserId, result1.UserId, database)
	chatId, _ := chat.InsertNewChat(result1.UserId, result2.UserId, database)
	message := structs.Message{
		ChatId:     chatId,
		SenderId:   result1.UserId,
		ReceiverId: result2.UserId,
		Content:    "Hello World",
	}
	t.Run("Test Insert Message", func(t *testing.T) {
		msg, err := messages.InsertMessage(message, *database)
		l.LogMessage("message_test.go", "Test Insert Message", msg)
		if err != nil {
			t.Errorf("Error inserting message: %v", err)
		}
	})
	t.Run("Test Get Messages", func(t *testing.T) {
		msgs, err := messages.GetPrivateMessages(chatId, *database)
		l.LogMessage("message_test.go", "Test Get Messages", msgs)
		if err != nil {
			t.Errorf("Error getting messages: %v", err)
		}
	})
}

func TestMessageHandler(t *testing.T) {
	Env := handler.Env{Env: database}
	email := "hello" + uuid.NewV4().String() + "@test.com"
	user1 := &structs.User{
		FirstName: "Adriell", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user1, *database)
	var result1 structs.User
	auth.GetUser("email", user1.Email, &result1, *database)
	email = "hello" + uuid.NewV4().String() + "@test.com"
	user2 := &structs.User{
		FirstName: "Nate", LastName: "LastTest", NickName: "NickTest", Email: email, Password: "Password123",
		DateOfBirth: "0001-01-01T00:00:00Z", AboutMe: "Test about me section", Avatar: "testPath", CreatedAt: "", UserId: "", SessionId: "-",
		IsLoggedIn: 0, IsPublic: 0, NumFollowers: 5, NumFollowing: 5, NumPosts: 0,
	}
	auth.InsertUser(*user2, *database)
	var result2 structs.User
	auth.GetUser("email", user2.Email, &result2, *database)
	follow.FollowUser(result1.UserId, result2.UserId, database)
	follow.FollowUser(result2.UserId, result1.UserId, database)
	chatId, _ := chat.InsertNewChat(result1.UserId, result2.UserId, database)
	message := structs.Message{
		ChatId:      chatId,
		MessageType: "privateMessage",
		SenderId:    result1.UserId,
		ReceiverId:  result2.UserId,
		Content:     "Hello World",
	}
	_, err := messages.InsertMessage(message, *database)
	t.Run("Test Insert Message", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error inserting message: %v", err)
		}
	})
	sampleUser := &structs.User{
		Email: email, Password: "Password123",
	}
	sampleUserBytes, err := json.Marshal(sampleUser)
	t.Run("Test Marshal data", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error marshalling the sampleUser")
		}
	})
	testReq := bytes.NewReader(sampleUserBytes) // Create the bytes into a reader
	req := httptest.NewRequest(http.MethodPost, "/login", testReq)
	w := httptest.NewRecorder()
	Env.Login(w, req)
	t.Run("Test login", func(t *testing.T) {
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	nr := httptest.NewRequest(http.MethodGet, "/message", nil)
	nr.Header = http.Header{"Cookie": w.Header()["Set-Cookie"]}
	values := nr.URL.Query()
	values.Add("chatId", chatId)
	nr.URL.RawQuery = values.Encode()
	nrr := httptest.NewRecorder()
	Env.Message(nrr, nr)
	t.Run("Test message handler status code", func(t *testing.T) {
		if nrr.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, nrr.Code)
		}
	})
	var result []structs.Message
	l.LogMessage("Test", "Messages", nrr.Body.Bytes())
	gotErr := json.Unmarshal(nrr.Body.Bytes(), &result)
	fmt.Print(result)
	t.Run("Test Message Handler", func(t *testing.T) {
		if gotErr != nil {
			t.Errorf("Error unmarshalling result: %v", gotErr)
			return
		}
	})
}
