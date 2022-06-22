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

    console.log(response);

    return {
      payload: res.data,
      status: res.code,
      message: res.message,
      error: null,
      success: true,
    };
  } catch (error) {
    console.log({ error });
    return {
      status: error.response.status,
      statusText: error.response.statusText,
      message: error.message,
      payload: error.response.data,
      success: false,
    };
  }
};

export const SERVICE_AUTH = { login };
