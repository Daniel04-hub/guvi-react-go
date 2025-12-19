import API_BASE_URL from '../config';
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useTheme } from '../ThemeContext';

const Profile = () => {
    const [profile, setProfile] = useState({
        email: '',
        age: '',
        dob: '',
        contact: '',
        address: ''
    });
    const [message, setMessage] = useState('');
    const navigate = useNavigate();
    const { theme, toggleTheme } = useTheme();

    const [countryCode, setCountryCode] = useState('+91');

    useEffect(() => {
        const fetchProfile = async () => {
            const token = localStorage.getItem('token');
            if (!token) {
                navigate('/login');
                return;
            }

            try {
                const res = await axios.get(`${API_BASE_URL}/profile`, {
                    headers: { Authorization: `Bearer ${token}` }
                });

                let contact = res.data.contact || '';
                setProfile(prev => ({ ...prev, ...res.data, contact }));

            } catch (err) {
                if (err.response?.status === 401) {
                    localStorage.removeItem('token');
                    navigate('/login');
                }
            }
        };
        fetchProfile();
    }, [navigate]);

    const handleChange = (e) => {
        setProfile({ ...profile, [e.target.name]: e.target.value });
    };

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    const validatePhoneNumber = (phone, code) => {
        const cleanPhone = phone.replace(/\D/g, '');
        if (code === '+91' || code === '+1') {
            return cleanPhone.length === 10;
        }
        return cleanPhone.length >= 7 && cleanPhone.length <= 15;
    };

    const validateAge = (dob, age) => {
        if (!dob || !age) return true; // Skip if empty
        const birthDate = new Date(dob);
        const today = new Date();
        let calculatedAge = today.getFullYear() - birthDate.getFullYear();
        const m = today.getMonth() - birthDate.getMonth();
        if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
            calculatedAge--;
        }
        return parseInt(age) === calculatedAge;
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setMessage(''); // Clear previous messages

        if (!validatePhoneNumber(profile.contact, countryCode)) {
            setMessage(`Invalid phone number for ${countryCode}. Must be 10 digits.`);
            return;
        }

        if (!validateAge(profile.dob, profile.age)) {
            setMessage('Age does not match Date of Birth.');
            return;
        }

        const token = localStorage.getItem('token');

        try {
            await axios.post(`${API_BASE_URL}/profile`, profile, {
                headers: { Authorization: `Bearer ${token}` }
            });
            setMessage('Profile updated successfully!');
            setTimeout(() => setMessage(''), 3000);
        } catch (err) {
            setMessage('Failed to update profile.');
        }
    };

    return (
        <div className="app-container" style={{ justifyContent: 'flex-start', paddingTop: '40px' }}>
            {/* Header Area */}
            <div style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                width: '100%',
                maxWidth: '600px',
                position: 'relative',
                marginBottom: '40px'
            }}>
                <h1 className="title" style={{ fontSize: '2rem', margin: 0 }}>NexEntry</h1>
                <button
                    onClick={handleLogout}
                    className="btn-outline-custom"
                    style={{
                        position: 'absolute',
                        right: 0,
                        top: '50%',
                        transform: 'translateY(-50%)'
                    }}
                >
                    Logout
                </button>
            </div>

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

                <div style={{ marginBottom: '2rem', textAlign: 'center' }}>
                    <h2>Details</h2>
                    <p className="subtitle" style={{ marginBottom: 0 }}>Manage your profile</p>
                </div>

                {message && <div style={{ color: 'var(--accent-color)', textAlign: 'center', marginBottom: '1rem' }}>{message}</div>}

                <form onSubmit={handleSubmit}>

                    <div style={{ display: 'flex', gap: '15px' }}>
                        <div className="form-group" style={{ flex: 1 }}>
                            <input
                                type="number"
                                name="age"
                                className="custom-input"
                                placeholder="Age"
                                value={profile.age}
                                onChange={handleChange}
                            />
                        </div>
                        <div className="form-group" style={{ flex: 1 }}>
                            <input
                                type="date"
                                name="dob"
                                className="custom-input"
                                value={profile.dob}
                                onChange={handleChange}
                                max="9999-12-31"
                            />
                        </div>
                    </div>

                    <div className="form-group">
                        <div style={{ display: 'flex', gap: '10px' }}>
                            <select
                                className="custom-input"
                                style={{ width: '110px' }}
                                value={countryCode}
                                onChange={(e) => setCountryCode(e.target.value)}
                            >
                                <option value="+91">IN +91</option>
                                <option value="+1">US +1</option>
                            </select>
                            <input
                                type="text"
                                name="contact"
                                className="custom-input"
                                placeholder="81234 56789"
                                value={profile.contact}
                                onChange={handleChange}
                                style={{ flex: 1 }}
                            />
                        </div>
                    </div>

                    <div className="form-group">
                        <textarea
                            name="address"
                            className="custom-input"
                            placeholder="123 Main St, City, Country"
                            rows="3"
                            value={profile.address}
                            onChange={handleChange}
                            style={{ resize: 'none' }}
                        ></textarea>
                    </div>

                    <button type="submit" className="btn-primary-custom" style={{ marginTop: '1rem' }}>
                        Save Changes
                    </button>
                </form>
            </div>
        </div>
    );
};

export default Profile;
