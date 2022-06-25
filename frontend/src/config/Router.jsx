import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { USER_ROLES } from '@utils/constants';
import { PATH } from './path';
import RoleSelection from '../pages/Auth/RoleSelection';
import Login from '../pages/Auth/Login';
import Register from '../pages/Auth/Register';
import Dashboard from '@pages/dashboard/Dashboard';
import Account from '@pages/dashboard/Account';
import NoPage from '@pages/NoPage';
import Class from '@pages/class/Class';

function Router() {
  return (
    <div>
      <Routes>
        <Route exact path={PATH.LOGIN} element={<RoleSelection />} />
        <Route exact path={PATH.REGISTER} element={<Navigate to={PATH.LOGIN} />} />
        <Route path={`${PATH.LOGIN}/${USER_ROLES.TEACHER.value}`} element={<Login role={USER_ROLES.TEACHER} />} />
        <Route path={`${PATH.LOGIN}/${USER_ROLES.PARENT.value}`} element={<Login role={USER_ROLES.PARENT} />} />
        <Route path={`${PATH.LOGIN}/${USER_ROLES.STUDENT.value}`} element={<Login role={USER_ROLES.STUDENT} />} />
        <Route
          path={`${PATH.REGISTER}/${USER_ROLES.TEACHER.value}`}
          element={<Register role={USER_ROLES.TEACHER} />}
        />
        <Route
          path={`${PATH.REGISTER}/${USER_ROLES.PARENT.value}`}
          element={<Register role={USER_ROLES.PARENT} />}
        />
        <Route
          path={`${PATH.REGISTER}/${USER_ROLES.STUDENT.value}`}
          element={<Register role={USER_ROLES.STUDENT} />}
        />
        <Route path={PATH.DASHBOARD} element={<Dashboard />} />
        <Route path={PATH.ACCOUNT} element={<Account />} />
        <Route path={PATH.CLASS + ":id"} element={<Class />} />
        <Route path={PATH.NOPAGE} element={<NoPage />} />
      </Routes>
    </div>
  );
}

export default Router;
