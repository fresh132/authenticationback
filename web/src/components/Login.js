import React, { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { Link, useNavigate } from 'react-router-dom';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [messageType, setMessageType] = useState('');

  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage('');

    if (!email || !password) {
      setMessage('Пожалуйста, заполните все поля');
      setMessageType('error');
      setLoading(false);
      return;
    }
    
    if (password.length < 8) {
      setMessage('Пароль должен содержать минимум 8 символов');
      setMessageType('error');
      setLoading(false);
      return;
    }

    const result = await login(email, password);
    
    if (result.success) {
      setMessage('Вход выполнен успешно!');
      setMessageType('success');
      setTimeout(() => {
        navigate('/welcome');
      }, 1000);
    } else {
      setMessage(result.error || 'Ошибка входа');
      setMessageType('error');
    }
    
    setLoading(false);
  };

  return (
    <div className="auth-container">
      <div className="auth-form">
        <h2>Вход в систему</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email:</label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              disabled={loading}
              placeholder="your@email.com"
              autoComplete="email"
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Пароль:</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              disabled={loading}
              minLength="8"
              placeholder="Минимум 8 символов"
              autoComplete="current-password"
            />
          </div>
          <button type="submit" className="btn" disabled={loading}>
            {loading ? 'Вход...' : 'Войти'}
          </button>
        </form>
        
        {message && (
          <div className={messageType === 'error' ? 'error-message' : 'success-message'}>
            {message}
          </div>
        )}
        
        <Link to="/register" className="link">
          Нет аккаунта? Зарегистрируйтесь
        </Link>
      </div>
    </div>
  );
};

export default Login;