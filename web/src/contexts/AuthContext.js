import React, { createContext, useContext, useState, useEffect } from 'react';
import { authService } from '../services/authService';

const AuthContext = createContext();

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(null);
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const initAuth = async () => {
      try {
        const savedToken = localStorage.getItem('token');
        if (savedToken) {
          setToken(savedToken);
          
          // Проверяем валидность токена
          try {
            const userData = await authService.getProfile(savedToken);
            // Бэкенд возвращает {user: {id, mail, created_at}}
            if (userData && userData.user) {
              setUser(userData.user);
            } else {
              throw new Error('Неверный формат данных');
            }
          } catch (error) {
            // Токен невалидный, удаляем его
            console.log('Token invalid, removing...', error.message);
            localStorage.removeItem('token');
            setToken(null);
          }
        }
      } catch (error) {
        console.error('Auth initialization error:', error);
      } finally {
        setLoading(false);
      }
    };

    initAuth();
  }, []);

  const login = async (email, password) => {
    try {
      const response = await authService.login(email, password);
      const { token } = response;
      setToken(token);
      localStorage.setItem('token', token);
      
      // Получаем данные пользователя
      try {
        const userData = await authService.getProfile(token);
        if (userData && userData.user) {
          setUser(userData.user);
        } else {
          console.error('Invalid user data format:', userData);
        }
      } catch (error) {
        console.error('Error getting user data:', error);
      }
      
      return { success: true };
    } catch (error) {
      return { success: false, error: error.message };
    }
  };

  const register = async (email, password) => {
    try {
      await authService.register(email, password);
      return { success: true };
    } catch (error) {
      return { success: false, error: error.message };
    }
  };

  const logout = () => {
    setToken(null);
    setUser(null);
    localStorage.removeItem('token');
  };

  const updatePassword = async (newPassword) => {
    try {
      await authService.changePassword(token, newPassword);
      return { success: true };
    } catch (error) {
      return { success: false, error: error.message };
    }
  };

  const deleteProfile = async () => {
    try {
      await authService.deleteProfile(token);
      logout();
      return { success: true };
    } catch (error) {
      return { success: false, error: error.message };
    }
  };

  const value = {
    token,
    user,
    login,
    register,
    logout,
    updatePassword,
    deleteProfile,
    loading
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};