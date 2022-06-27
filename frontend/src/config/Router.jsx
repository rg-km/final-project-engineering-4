import React from 'react';
import { Navigate, Route, Routes, Outlet } from 'react-router-dom';
import { USER_ROLES } from '@utils/constants';
import { PATH } from './path';
import RoleSelection from '../pages/Auth/RoleSelection';
import Login from '../pages/Auth/Login';
import Register from '../pages/Auth/Register';
import Dashboard from '@pages/DashboardPage/Dashboard';
import Account from '@pages/DashboardPage/Account';
import NoPage from '@pages/NoPage';
import { useAuthStore } from '@store/auth.store';

const PATH_TEACHER = USER_ROLES.TEACHER.value;
const PATH_PARENT = USER_ROLES.PARENT.value;
const PATH_STUDENT = USER_ROLES.STUDENT.value;

function AuthenticationRoute() {
  const { token } = useAuthStore((state) => state);
  return token ? <Navigate to={PATH.DASHBOARD} /> : <Outlet />;
}

function ProtectedRoute() {
  const { token } = useAuthStore((state) => state);
  return token ? <Outlet /> : <Navigate to={PATH.LOGIN} />;
}

function Router() {
  return (
    <div>
      <Routes>
        <Route element={<AuthenticationRoute />}>
          <Route exact path={PATH.LOGIN} element={<RoleSelection />} />
          <Route exact path={PATH.REGISTER} element={<Navigate to={PATH.LOGIN} />} />
          <Route path={`${PATH.LOGIN}/${PATH_PARENT}`} element={<Login role={USER_ROLES.PARENT} />} />
          <Route path={`${PATH.LOGIN}/${PATH_STUDENT}`} element={<Login role={USER_ROLES.STUDENT} />} />
          <Route path={`${PATH.LOGIN}/${PATH_TEACHER}`} element={<Login role={USER_ROLES.TEACHER} />} />
          <Route path={`${PATH.REGISTER}/${PATH_TEACHER}`} element={<Register role={USER_ROLES.TEACHER} />} />
          <Route path={`${PATH.REGISTER}/${PATH_PARENT}`} element={<Register role={USER_ROLES.PARENT} />} />
          <Route path={`${PATH.REGISTER}/${PATH_STUDENT}`} element={<Register role={USER_ROLES.STUDENT} />} />
        </Route>

        <Route element={<ProtectedRoute />}>
          <Route path={PATH.DASHBOARD} element={<Dashboard />} />
          <Route path={PATH.ACCOUNT} element={<Account />} />
        </Route>

        <Route path={PATH.NOPAGE} element={<NoPage />} />
      </Routes>
    </div>
  );
}

export default Router;
