import React from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

const Welcome = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleProfileClick = () => {
    navigate('/profile');
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <div className="auth-container">
      <div className="welcome-container">
        <h1>Добро пожаловать! 👋</h1>
        <p>
          {user?.mail ? `Рады видеть вас снова, ${user.mail}!` : 'Вы успешно вошли в систему.'}
        </p>
        <p>
          Нажмите на кнопку ниже, чтобы посмотреть информацию о вашем аккаунте.
        </p>
        
        {/* Отладочная информация */}
        {process.env.NODE_ENV === 'development' && user && (
          <div style={{ 
            margin: '20px 0', 
            padding: '15px', 
            backgroundColor: '#f5f5f5', 
            borderRadius: '8px',
            fontSize: '12px',
            textAlign: 'left',
            border: '1px solid #ddd'
          }}>
            <strong>Отладочная информация:</strong>
            <br />ID: {user.id || 'Нет'}
            <br />Email: {user.mail || 'Нет'}
            <br />Дата: {user.created_at || 'Нет'}
          </div>
        )}
        
        <button 
          className="messenger-button"
          onClick={handleProfileClick}
        >
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
            <circle cx="12" cy="7" r="4"></circle>
          </svg>
          Мой профиль
        </button>
        
        <button 
          className="btn btn-secondary" 
          onClick={handleLogout}
          style={{ marginTop: '20px' }}
        >
          Выйти
        </button>
      </div>
    </div>
  );
};

export default Welcome;