import Home from '../Home';
import Login from '../Login';
import Chat from '../Chat';
import { Route, Routes } from 'react-router';
import PrivateRoutes from '../PrivateRoutes';

const RouterWrapper = () => {
    return (
        <Routes>
            <Route element={<PrivateRoutes />}>
                <Route index element={<Home />} />
                <Route path='chat' element={<Chat />} />
            </Route>
            <Route path='login' element={<Login />} />
            <Route path='*' element={<>Not Found</>} />
        </Routes>
    );
};

export default RouterWrapper;
