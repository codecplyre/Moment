import './App.css';
import Footer from './Layouts/Footer/Footer';
import Header from './Layouts/Header/Header';
import Home from './Pages/Home/Index';
import { Route, Routes } from 'react-router-dom';
import Login from './Pages/LoginPage/Login';
import Registration from './Pages/RegPage/Registration';
import Chat from './Features/Chat/Chat';
import Profile from './Pages/Profile/Profile';
import ProfileInfoPopUp from './Features/Profile/ProfileInfoPopUp';
import Comments from './Features/Comments/Index';
import { useEffect, useState } from 'react';
import NewPost from './Features/Newpost/NewPost';
import { Notification } from './Features/Notification/Notification';
import Validation from './Components/Validation/Validation';
import { Menu } from './Layouts/Menu/Menu';
import ValidRedirect from './Components/Validation/ValidRedirect';
import useWindowDimensions from './Components/Hooks/UseWindowDimensions';
import CloseFriendsUsers from './Features/Profile/CloseFriendsUsers';
import Followers from './Features/Profile/Followers';
import Following from './Features/Profile/Following';
import { SearchModal } from './Features/Search/SearchModal';
function App() {
    const [authorised, setAuthorised] = useState(false);
    Validation(setAuthorised);
    const [socket, setSocket] = useState(null);
    const [isMenuOpen, setIsMenuOpen] = useState(false);
    const [isSearchModalOpen, setIsSearchModalOpen] = useState(false);
    const { width } = useWindowDimensions();
    const [query, setQuery] = useState('');
    let isMobile = width < 600;
    return (
        <div
            className='App'
            onClick={() => {
                setIsMenuOpen(false);
                setIsSearchModalOpen(false);
            }}>
            {authorised && (
                <Header
                    setSocket={setSocket}
                    setIsMenuOpen={setIsMenuOpen}
                    setIsSearchModalOpen={setIsSearchModalOpen}
                    onChange={(e) => {
                        setQuery(e.target.value);
                    }}
                />
            )}
            <>
                {isMenuOpen ? (
                    <Menu setIsMenuOpen={setIsMenuOpen} auth={setAuthorised} />
                ) : null}
            </>
            <>{isSearchModalOpen ? <SearchModal query={query} /> : null}</>
            <>
                <Routes>
                    <Route
                        path='/'
                        element={
                            <ValidRedirect>
                                <Login auth={setAuthorised} />
                            </ValidRedirect>
                        }
                    />
                    <Route
                        path='/register'
                        element={
                            <ValidRedirect>
                                <Registration />
                            </ValidRedirect>
                        }
                    />
                    {/* need to be replaced */}
                    <Route path='*' element={<></>} />
                </Routes>
            </>
            {authorised && (
                <>
                    <Routes>
                        <Route
                            path='/home'
                            element={<Home isMobile={isMobile} />}
                        />
                        <Route path='/newpost' element={<NewPost />} />
                        <Route
                            path='/messages/*'
                            element={
                                <Chat isMobile={isMobile} socket={socket} />
                            }
                        />
                        <Route path='/groups' element={<h1>Groups</h1>} />
                        <Route
                            path='/comments'
                            element={<Comments isMobile={isMobile} />}
                        />
                        <Route path='notifications' element={<></>} />
                        <Route
                            path='notifications/:type'
                            element={<Notification socket={socket} />}
                        />
                        <Route path='/profile' element={<Profile />} />
                        <Route
                            path='/closefriends'
                            element={<CloseFriendsUsers />}
                        />
                        <Route path='/followers' element={<Followers />} />
                        <Route path='/following' element={<Following />} />
                        <Route
                            path='/update'
                            element={<ProfileInfoPopUp styleName='popUp' />}
                        />
                    </Routes>
                </>
            )}
            {authorised ? (
                <Footer setIsSearchModalOpen={setIsSearchModalOpen} />
            ) : null}
        </div>
    );
}

export default App;
