import { Navigate, Outlet } from 'react-router';

const PrivateRoutes = () => {
    const isAuthenticated = false;
    console.log('isAuthenticated', isAuthenticated);
    return isAuthenticated ? <Outlet /> : <Navigate to='/login' />;
};
export default PrivateRoutes;
