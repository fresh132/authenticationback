import React, { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

const Profile = () => {
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('');
  const [showPasswordForm, setShowPasswordForm] = useState(false);

  const { user, updatePassword, deleteProfile, logout } = useAuth();
  const navigate = useNavigate();

  const handlePasswordUpdate = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage('');

    if (newPassword !== confirmPassword) {
      setMessage('Пароли не совпадают');
      setMessageType('error');
      setLoading(false);
      return;
    }

    if (newPassword.length < 8) {
      setMessage('Пароль должен содержать минимум 8 символов');
      setMessageType('error');
      setLoading(false);
      return;
    }

    const result = await updatePassword(newPassword);
    
    if (result.success) {
      setMessage('Пароль успешно обновлен');
      setMessageType('success');
      setNewPassword('');
      setConfirmPassword('');
      setShowPasswordForm(false);
    } else {
      setMessage(result.error || 'Ошибка при обновлении пароля');
      setMessageType('error');
    }
    
    setLoading(false);
  };

  const handleDeleteProfile = async () => {
    if (window.confirm('Вы уверены, что хотите удалить свой профиль? Это действие необратимо.')) {
      setLoading(true);
      const result = await deleteProfile();
      
      if (!result.success) {
        setMessage(result.error || 'Ошибка при удалении профиля');
        setMessageType('error');
        setLoading(false);
      }
    }
  };

  const handleBack = () => {
    navigate('/welcome');
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'Неизвестно';
    try {
      const date = new Date(dateString);
      if (isNaN(date.getTime())) return 'Некорректная дата';
      return date.toLocaleString('ru-RU', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      });
    } catch (error) {
      return 'Ошибка даты';
    }
  };

  return (
    <div className="auth-container">
      <div className="profile-container">
        <div className="profile-header">
          <h1>Мой профиль</h1>
          <button className="back-button" onClick={handleBack}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            Назад
          </button>
        </div>

        <div className="profile-info">
          <h3>Информация об аккаунте</h3>
          {user ? (
            <>
              <div className="info-item">
                <label>ID пользователя:</label>
                <span>{user.id || 'Не указано'}</span>
              </div>
              <div className="info-item">
                <label>Email:</label>
                <span>{user.mail || 'Не указано'}</span>
              </div>
              <div className="info-item">
                <label>Дата регистрации:</label>
                <span>{formatDate(user.created_at)}</span>
              </div>
            </>
          ) : (
            <div style={{ 
              padding: '20px', 
              textAlign: 'center', 
              color: '#666',
              backgroundColor: '#FAFAFA',
              borderRadius: '10px',
              border: '1px solid #E3F2FD'
            }}>
              <p>Загрузка данных профиля...</p>
              <p style={{ fontSize: '14px', marginTop: '10px' }}>
                Если данные не загружаются долго, попробуйте обновить страницу
              </p>
            </div>
          )}
        </div>

        {showPasswordForm && (
          <div style={{ 
            padding: '20px', 
            backgroundColor: '#FAFAFA', 
            borderRadius: '10px',
            marginBottom: '20px',
            border: '1px solid #E3F2FD'
          }}>
            <h3 style={{ color: '#1976D2', marginBottom: '20px' }}>Сменить пароль</h3>
            <form onSubmit={handlePasswordUpdate}>
              <div className="form-group">
                <label htmlFor="newPassword">Новый пароль:</label>
                <input
                  type="password"
                  id="newPassword"
                  value={newPassword}
                  onChange={(e) => setNewPassword(e.target.value)}
                  required
                  disabled={loading}
                  minLength="8"
                  placeholder="Минимум 8 символов"
                />
              </div>
              <div className="form-group">
                <label htmlFor="confirmPassword">Подтвердите новый пароль:</label>
                <input
                  type="password"
                  id="confirmPassword"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  required
                  disabled={loading}
                  minLength="8"
                  placeholder="Повторите пароль"
                />
              </div>
              <div style={{ display: 'flex', gap: '10px' }}>
                <button type="submit" className="btn" disabled={loading}>
                  {loading ? 'Обновление...' : 'Обновить пароль'}
                </button>
                <button 
                  type="button" 
                  className="btn btn-secondary" 
                  onClick={() => {
                    setShowPasswordForm(false);
                    setNewPassword('');
                    setConfirmPassword('');
                    setMessage('');
                  }}
                  disabled={loading}
                >
                  Отмена
                </button>
              </div>
            </form>
          </div>
        )}

        {message && (
          <div className={messageType === 'error' ? 'error-message' : 'success-message'}>
            {message}
          </div>
        )}

        <div className="action-buttons">
          <button 
            className="btn btn-secondary" 
            onClick={() => setShowPasswordForm(!showPasswordForm)}
            disabled={loading || !user}
          >
            {showPasswordForm ? 'Отменить' : 'Изменить пароль'}
          </button>
          <button 
            className="btn btn-danger" 
            onClick={handleDeleteProfile}
            disabled={loading || !user}
          >
            Удалить профиль
          </button>
        </div>
      </div>
    </div>
  );
};

export default Profile;