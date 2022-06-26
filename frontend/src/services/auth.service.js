import { USER_ROLES } from '@utils/constants';
import { http } from './http';

const login = async (role, email, password) => {
  let baseURL = null;

  const data = {
    email,
    password,
  };

  if (role === USER_ROLES.TEACHER.value) baseURL = '/guru/login';
  if (role === USER_ROLES.PARENT.value) baseURL = '/orangtua/login';
  if (role === USER_ROLES.STUDENT.value) baseURL = '/siswa/login';

  try {
    const response = await http.post(baseURL, data);

    return {
      payload: response.data,
      status: response.code,
      message: response.message,
      success: true,
    };
  } catch (error) {
    return {
      payload: error.response.data,
      status: error.response.status,
      message: error?.response?.data?.message || error.message,
      success: false,
    };
  }
};

export const SERVICE_AUTH = { login };
