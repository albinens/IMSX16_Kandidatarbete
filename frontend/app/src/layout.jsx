import React, {useState, useEffect} from 'react'
import {Outlet, Link} from 'react-router-dom'


import MobileNav from './components/mobileNav/mobileNav'
import DesktopNav from './components/desktopNav/desktopNav'

const Layout = () => {

  let mobileWidth = 700;
  const [windowWidth, setWindowWidth] = useState(window.innerWidth);

  useEffect(() => {
    const updateWindowDimensions = () => {
      const newWindowWidth = window.innerWidth;
      setWindowWidth(newWindowWidth);
      console.log("updating width");
    };

    window.addEventListener("resize", updateWindowDimensions);

    return () => window.removeEventListener("resize", updateWindowDimensions) 

  }, []);
  return (
    <>
      {windowWidth >= mobileWidth ? <DesktopNav /> : null}
      <Outlet />
      {windowWidth < mobileWidth ? <MobileNav /> : null}
    </>
  )
}

export default Layout
