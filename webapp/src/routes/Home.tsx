import * as React from "react";
import request from "../utils/request";

interface IUserInfo {
  id: number;
  name: string;
  avatar_url: string;
}
const Home = ():JSX.Element => {
  const [userInfo, setUserInfo] = React.useState({} as Partial<IUserInfo>);

  React.useEffect(() => {
    request({
      method: "POST",
      action: "/api/login"
    }).then((res: any) => {
      if (res.RetCode === 0 && res.Data) {
        setUserInfo(res.Data[0])
      }
    })
  }, []);

  return (
      <div className="login-div">
        <span>{userInfo.name}</span>
        <img src={userInfo.avatar_url} alt=""/>
      </div>
  )
};

export default Home;
