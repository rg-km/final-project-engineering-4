import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { USER_ROLES } from '@utils/constants';
import { PATH } from './path';
import RoleSelection from './Auth/RoleSelection';
import Login from './Auth/Login';
import Register from './Auth/Register';

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
      </Routes>
    </div>
  );
}

export default Router;
