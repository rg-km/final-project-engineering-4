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
      success: true,
      message: response.data.message,
      status: response.data.code,
      payload: response.data.data,
    };
  } catch (error) {
    return {
      success: false,
      message: error?.response?.data?.message || error.message,
      status: error.response.status,
      payload: error.response.data,
    };
  }
};

export const SERVICE_AUTH = { login };
