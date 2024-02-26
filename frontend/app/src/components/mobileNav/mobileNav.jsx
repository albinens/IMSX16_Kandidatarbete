import React from 'react';
import './mobileNav.css';
import search from '../../assets/search.svg';
import settings from '../../assets/settings.svg';
import home from '../../assets/home.svg';

const MobileNav = () => {
  return (
    <nav className='mobile-nav'>
      <ul className='mobile-nav-list'>
        <li><a href="/"><img src={home}/></a></li>
        <li><a href="/search"><img src={search}/></a></li>
        <li><a href="/settings"><img src={settings}/></a></li>
      </ul>
    </nav>
  );
}

export default MobileNav;
