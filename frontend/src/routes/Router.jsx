import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import RoleSelection from './Auth/RoleSelection';
import Login from './Auth/Login';
import { PATH } from './path';
import { USER_ROLES } from '@utils/constants';

function Router() {
  return (
    <div>
      <Routes>
        <Route exact path={PATH.LOGIN} element={<RoleSelection />} />
        <Route exact path={PATH.REGISTER} element={<Navigate to={PATH.LOGIN} />} />
        <Route path={`${PATH.LOGIN}/${USER_ROLES.TEACHER.value}`} element={<Login role={USER_ROLES.TEACHER} />} />
      </Routes>
    </div>
  );
}

export default Router;
