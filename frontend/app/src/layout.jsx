import React from 'react'
import {Outlet, Link} from 'react-router-dom'

import ResponsiveAppBar from './components/mui/responsiveAppBar/responsiveAppBar'

const Layout = () => {

  return (
    <>
      <ResponsiveAppBar /> 
      <Outlet />
    </>
  )
}

export default Layout
