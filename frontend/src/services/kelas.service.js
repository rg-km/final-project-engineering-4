import { http } from './http';
import { useAuthStore } from '@store/auth.store';

const { token } = useAuthStore((state) => state);

const headerConfig = { headers: { Authorization: `Bearer ${token}` } };

const getListKelas = async () => {
  try {
    const response = await http.get(baseURL, headerConfig);

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

export const SERVICE_KELAS = { getListKelas };
