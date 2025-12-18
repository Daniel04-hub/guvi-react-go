import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, Link } from 'react-router-dom';
import { useTheme } from '../ThemeContext';
import API_BASE_URL from '../config';

const Register = () => {
    const [formData, setFormData] = useState({ username: '', email: '', password: '' });
    const [error, setError] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const [passwordStrength, setPasswordStrength] = useState('');
    const navigate = useNavigate();
    const { theme, toggleTheme } = useTheme();

    const checkStrength = (pass) => {
        if (!pass) return '';
        if (pass.length < 6) return 'Weak';
        if (pass.match(/[a-zA-Z]/) && pass.match(/[0-9]/) && pass.length >= 8) {
            return 'Strong';
        }
        return 'Medium';
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });

        if (name === 'password') {
            const strength = checkStrength(value);
            setPasswordStrength(strength);
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        if (passwordStrength === 'Weak') {
            setError('Password is too weak.');
            return;
        }

        try {
            await axios.post(`${API_BASE_URL}/register`, formData);
            alert('Registration successful');
            navigate('/login');
        } catch (err) {
            setError(err.response?.data?.message || 'Registration failed');
        }
    };

    const getStrengthClass = () => {
        if (passwordStrength === 'Weak') return 'strength-weak';
        if (passwordStrength === 'Medium') return 'strength-medium';
        if (passwordStrength === 'Strong') return 'strength-strong';
        return '';
    };

    return (
        <div className="app-container">
            <div className="custom-card">
                <button className="theme-toggle" onClick={toggleTheme} title={theme === 'dark' ? 'Switch to Light Mode' : 'Switch to Dark Mode'}>
                    {theme === 'dark' ? (
                        <svg viewBox="0 0 24 24" fill="currentColor" stroke="none">
                            <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"></path>
                        </svg>
                    ) : (
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                            <circle cx="12" cy="12" r="5"></circle>
                            <line x1="12" y1="1" x2="12" y2="3"></line>
                            <line x1="12" y1="21" x2="12" y2="23"></line>
                            <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                            <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                            <line x1="1" y1="12" x2="3" y2="12"></line>
                            <line x1="21" y1="12" x2="23" y2="12"></line>
                            <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                            <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
                        </svg>
                    )}
                </button>

                <h2>Create Account</h2>
                <p className="subtitle">Sign up to get started</p>

                {error && <div style={{ color: '#ff4d4d', textAlign: 'center', marginBottom: '1rem', background: 'rgba(255, 77, 77, 0.1)', padding: '10px', borderRadius: '12px' }}>{error}</div>}

                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <input
                            type="text"
                            name="username"
                            className="custom-input"
                            placeholder="Username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    <div className="form-group">
                        <input
                            type="email"
                            name="email"
                            className="custom-input"
                            placeholder="name@example.com"
                            value={formData.email}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    <div className="form-group" style={{ position: 'relative' }}>
                        <input
                            type={showPassword ? "text" : "password"}
                            name="password"
                            className="custom-input"
                            placeholder="Password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                        />
                        <button
                            type="button"
                            className="toggle-password"
                            onClick={() => setShowPassword(!showPassword)}
                        >
                            {showPassword ? (
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
                            ) : (
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
                            )}
                        </button>
                        {formData.password && (
                            <div className={`strength-text ${getStrengthClass()}`}>
                                {passwordStrength}
                            </div>
                        )}
                    </div>

                    <button type="submit" className="btn-primary-custom" style={{ marginTop: '1.5rem' }}>
                        Sign Up <span style={{ fontSize: '1.2rem', marginLeft: '5px' }}>→</span>
                    </button>
                </form>

                <div style={{ textAlign: 'center', marginTop: '2rem' }}>
                    <span style={{ color: 'var(--text-muted)' }}>Already have an account? </span>
                    <Link to="/login" className="link-custom">Login</Link>
                </div>
            </div>
        </div>
    );
};

export default Register;
