import Home from '../Home';
import Login from '../Login';
import Chat from '../Chat';
import { Route, Routes } from 'react-router';

const RouterWrapper = () => {
    return (
        <Routes>
            <Route index element={<Home />} />
            <Route path='login' element={<Login />} />
            <Route path='chat' element={<Chat />} />
        </Routes>
    );
};

export default RouterWrapper;
