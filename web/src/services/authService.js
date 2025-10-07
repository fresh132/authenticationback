import axios from 'axios';

const api = axios.create({
  baseURL: '', 
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000, // Таймаут 10 секунд
});


api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);


api.interceptors.response.use(
  (response) => {

    if (response.config.url?.includes('/auth/user')) {
      console.log('Profile response:', response.data);
    }
    return response;
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export const authService = {
  login: async (email, password) => {
    try {
      const response = await api.post('/enter', { mail: email, password });
      console.log('Login response:', response.data);
      return response.data;
    } catch (error) {
      console.error('Login error:', error);
      const message = error.response?.data?.error || error.response?.data?.message || 'Ошибка входа';
      throw new Error(message);
    }
  },

  register: async (email, password) => {
    try {
      const response = await api.post('/register', { mail: email, password });
      console.log('Register response:', response.data);
      return response.data;
    } catch (error) {
      console.error('Register error:', error);
      const message = error.response?.data?.error || error.response?.data?.message || 'Ошибка регистрации';
      throw new Error(message);
    }
  },

  getProfile: async (token) => {
    try {
      const response = await api.get('/auth/user', {
        headers: { Authorization: `Bearer ${token}` }
      });
      console.log('Get profile response:', response.data);
      
  
      if (!response.data) {
        throw new Error('Пустой ответ от сервера');
      }
      
      return response.data;
    } catch (error) {
      console.error('Get profile error:', error);
      const message = error.response?.data?.error || error.response?.data?.message || 'Ошибка получения профиля';
      throw new Error(message);
    }
  },

  changePassword: async (token, newPassword) => {
    try {
      const response = await api.put('/auth/update', { newpassword: newPassword }, {
        headers: { Authorization: `Bearer ${token}` }
      });
      console.log('Change password response:', response.data);
      return response.data;
    } catch (error) {
      console.error('Change password error:', error);
      const message = error.response?.data?.error || error.response?.data?.message || 'Ошибка смены пароля';
      throw new Error(message);
    }
  },

  deleteProfile: async (token) => {
    try {
      const response = await api.delete('/auth/delete', {
        headers: { Authorization: `Bearer ${token}` }
      });
      console.log('Delete profile response:', response.data);
      return response.data;
    } catch (error) {
      console.error('Delete profile error:', error);
      const message = error.response?.data?.error || error.response?.data?.message || 'Ошибка удаления профиля';
      throw new Error(message);
    }
  }
};