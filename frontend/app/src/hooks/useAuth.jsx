import React from 'react';

function useAuth() {

  //key: 
  const [authenticated, setAuthenticated] = React.useState(localStorage.getItem("authed") === 'true' ? true : false)
  const [authKey, setAuthKey] = React.useState('7xD9MFMRT8WstAYhrP88pEGd9kpRHq69')

  return { authenticated, setAuthenticated, authKey, setAuthKey }

}

export default useAuth;