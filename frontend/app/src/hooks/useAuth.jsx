import React from 'react';
import axios from 'axios';

function useAuth() {

  //key: 7xD9MFMRT8WstAYhrP88pEGd9kpRHq69

  const [authenticated, setAuthenticated] = React.useState(localStorage.getItem("authed") === 'true' ? true : false)
  const [authKey, setAuthKey] = React.useState('')

  const handleAuthKeySet = (key) => {
    //Kolla backend om nyckel är giltig
    setAuthKey(key)
    setAuthenticated(true)
    localStorage.setItem("authed", 'true')
  }


  return { authenticated, setAuthenticated, authKey, setAuthKey }

}

export default useAuth;