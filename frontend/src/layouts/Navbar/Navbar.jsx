import { NavLink } from 'react-router-dom';
import {
    HomeIcon,
    SearchIcon,
    NewPostIcon,
    MessagesIcon,
    GroupsIcon,
} from '../../components/Icons/Icons';
import './Navbar.css';

const FooterNav = () => {
    return (
        <div className='navbar'>
            {/* Font Awesome Pro 6.2.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2022 Fonticons, Inc.*/}
            <>
                <NavLink to='/home'>
                    <HomeIcon />
                </NavLink>
                <NavLink to='/search'>
                    <SearchIcon />
                </NavLink>
                <NavLink to='/newpost'>
                    <NewPostIcon />
                </NavLink>
                <NavLink to='/messages'>
                    <MessagesIcon />
                </NavLink>
                <NavLink to='/groups'>
                    <GroupsIcon />
                </NavLink>
            </>
        </div>
    );
};

export default FooterNav;
