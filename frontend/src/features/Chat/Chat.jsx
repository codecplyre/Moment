import './Chat.css';
import Body from '../../components/Body/Body';
import Card from '../../components/card/Card';
import MiniUserCard from '../../components/MiniUserCard/MiniUserCard';
import SendMessageBox from './components/SendMessageBox';
const Chat = ({ bodyStyleName, cardStyleName }) => {
    return (
        <Body styleName={bodyStyleName}>
            <Card styleName={cardStyleName}>
                <div className='chatContainer'>
                    <div className='chatBox'>
                        <div className='chatUsersContainer'>
                            <div className='chatContainerHeader'>
                                <div>
                                    <div className='userName'>Moment</div>
                                </div>
                                <div>
                                    <svg
                                        height='22'
                                        width='22'
                                        xmlns='http://www.w3.org/2000/svg'
                                        viewBox='0 0 512 512'>
                                        <path d='M373.1 24.97C401.2-3.147 446.8-3.147 474.9 24.97L487 37.09C515.1 65.21 515.1 110.8 487 138.9L289.8 336.2C281.1 344.8 270.4 351.1 258.6 354.5L158.6 383.1C150.2 385.5 141.2 383.1 135 376.1C128.9 370.8 126.5 361.8 128.9 353.4L157.5 253.4C160.9 241.6 167.2 230.9 175.8 222.2L373.1 24.97zM440.1 58.91C431.6 49.54 416.4 49.54 407 58.91L377.9 88L424 134.1L453.1 104.1C462.5 95.6 462.5 80.4 453.1 71.03L440.1 58.91zM203.7 266.6L186.9 325.1L245.4 308.3C249.4 307.2 252.9 305.1 255.8 302.2L390.1 168L344 121.9L209.8 256.2C206.9 259.1 204.8 262.6 203.7 266.6zM200 64C213.3 64 224 74.75 224 88C224 101.3 213.3 112 200 112H88C65.91 112 48 129.9 48 152V424C48 446.1 65.91 464 88 464H360C382.1 464 400 446.1 400 424V312C400 298.7 410.7 288 424 288C437.3 288 448 298.7 448 312V424C448 472.6 408.6 512 360 512H88C39.4 512 0 472.6 0 424V152C0 103.4 39.4 64 88 64H200z' />
                                    </svg>
                                </div>
                            </div>
                            <div className='chatUsers scrollbar-hidden'>
                                <MiniUserCard
                                    img='./logo.svg'
                                    propsId='1'
                                    name='Moment'
                                    content='online'
                                />
                                <MiniUserCard
                                    img='./logo.svg'
                                    propsId='1'
                                    name='Moment'
                                    content='online'
                                />
                                <MiniUserCard
                                    img='./logo.svg'
                                    propsId='1'
                                    name='Moment'
                                    content='online'
                                />
                                <MiniUserCard
                                    img='./logo.svg'
                                    propsId='1'
                                    name='Ricky'
                                    content='online'
                                />
                            </div>
                        </div>
                        <div className='sendMessageContainer'>
                            <SendMessageBox />
                        </div>
                    </div>
                </div>
            </Card>
        </Body>
    );
};
export default Chat;
