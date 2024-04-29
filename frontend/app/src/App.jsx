import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Layout from './layout'
import DataBoard from './pages/dataBoard'
import ListRooms from './pages/listRooms'
import Sensors from './pages/sensors'
import Settings from './pages/settings'
import About from './pages/about'

function App() {
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route path="/" element={<ListRooms />}/>
          <Route path="listRooms" element={<ListRooms />}/>
          <Route path="settings" element={<Settings />}/>
          <Route path='sensors' element={<Sensors />}/>
          <Route path='databoard' element={<DataBoard />}/>
          <Route path='about' element={<About />}/>
          <Route path='*' element={<h1>Not Found</h1>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )

}

export default App
