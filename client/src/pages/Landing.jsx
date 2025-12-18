import React from 'react';
import { Link } from 'react-router-dom';
import { useTheme } from '../ThemeContext';

const Landing = () => {
    const { theme, toggleTheme } = useTheme();

    return (
        <div className="app-container" style={{ position: 'relative', overflow: 'hidden' }}>
            {/* Theme Toggle Top Right */}
            <button
                className="theme-toggle"
                onClick={toggleTheme}
                title="Toggle Theme"
                style={{ top: 32, right: 32, position: 'absolute' }}
            >
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

            {/* Content directly on background, no card */}
            <div className="landing-content" style={{ textAlign: 'center', zIndex: 1, position: 'relative' }}>
                <h1 className="title" style={{
                    fontSize: '4rem',
                    marginBottom: '1rem',
                    lineHeight: '1.1',
                    letterSpacing: '-0.03em'
                }}>
                    NexEntry
                </h1>
                <p className="subtitle" style={{
                    fontSize: '1.2rem',
                    marginBottom: '3.5rem',
                    opacity: 0.7,
                    fontWeight: 400
                }}>
                    A Full-Stack Application with MySQL, MongoDB, and Redis
                </p>

                <div style={{ display: 'flex', justifyContent: 'center' }}>
                    <Link to="/register" className="btn-primary-custom" style={{
                        padding: '18px 48px',
                        fontSize: '1.1rem',
                        width: 'auto',  /* Allow button to size to content */
                        minWidth: '200px'
                    }}>
                        Let's Start <span style={{ fontSize: '1.3rem', marginLeft: '10px' }}>→</span>
                    </Link>
                </div>
            </div>
        </div>
    );
};

export default Landing;
