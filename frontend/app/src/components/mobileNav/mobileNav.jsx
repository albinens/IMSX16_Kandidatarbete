import React from 'react';
import './mobileNav.css';

const MobileNav = () => {
  return (
    <nav>
      <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/search">Search</a></li>
        <li><a href="/settings">Settings</a></li>
      </ul>
    </nav>
  );
}

export default MobileNav;
