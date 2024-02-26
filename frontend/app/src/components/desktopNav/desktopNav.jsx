import React from 'react';
import './desktopNav.css';

const DesktopNav = () => {
  return (
    <nav className="desktop-nav">
      <ul className="desktop-nav-list">
        <li><a href="/">Home</a></li>
        <li><a href="/search">Search</a></li>
        <li><a href="/settings">Settings</a></li>
      </ul>
    </nav>
  );
}

export default DesktopNav;