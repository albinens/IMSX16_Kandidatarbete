import React, { useEffect } from 'react';

function useAuth() {

  //key: 7xD9MFMRT8WstAYhrP88pEGd9kpRHq69


  const [authenticated, setAuthenticated] = React.useState(localStorage.getItem("authed") === 'true' ? true : false)
  const [authKey, setAuthKey] = React.useState('')

  const handleAuthKeySet = (key) => {
    console.log("HÄR")
    if(key === '7xD9MFMRT8WstAYhrP88pEGd9kpRHq69'){
      console.log('Key is valid')
      //Kolla backend om nyckel är giltig
      setAuthKey(key)
      setAuthenticated(true)
      localStorage.setItem("authed", true)
      localStorage.setItem("authKey", key)
    }
  }


  return { authenticated, setAuthenticated, authKey, handleAuthKeySet }

}

export default useAuth;