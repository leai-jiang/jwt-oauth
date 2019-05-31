import * as React from "react";

const Login = ():JSX.Element => {

  return (
      <div className="login-div">
          <ul>
              <li>
                  <a href="/api/login/github">
                      <i className="iconfont github"/>
                      <span>Github</span>
                  </a>
              </li>
          </ul>
      </div>
  )
};

export default Login;
