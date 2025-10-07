import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import { useAuth } from './contexts/AuthContext';
import Login from './components/Login';
import Register from './components/Register';
import Welcome from './components/Welcome';
import Profile from './components/Profile';
import './App.css';

function AppRoutes() {
  const { token, loading } = useAuth();

 
  if (loading) {
    return (
      <div className="loading-container">
        <div className="loading-spinner"></div>
        <p>Загрузка...</p>
      </div>
    );
  }

  return (
    <div className="App">
      <Routes>
        <Route 
          path="/login" 
          element={!token ? <Login /> : <Navigate to="/welcome" replace />} 
        />
        <Route 
          path="/register" 
          element={!token ? <Register /> : <Navigate to="/welcome" replace />} 
        />
        <Route 
          path="/welcome" 
          element={token ? <Welcome /> : <Navigate to="/login" replace />} 
        />
        <Route 
          path="/profile" 
          element={token ? <Profile /> : <Navigate to="/login" replace />} 
        />
        <Route 
          path="/" 
          element={<Navigate to={token ? "/welcome" : "/login"} replace />} 
        />
      </Routes>
    </div>
  );
}

function App() {
  return (
    <AuthProvider>
      <Router>
        <AppRoutes />
      </Router>
    </AuthProvider>
  );
}

export default App;