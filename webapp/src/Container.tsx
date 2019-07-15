import * as React from "react";
import Background from "./background/Background";
import { Header } from "./layout"
import request from "./utils/request";
import GlobalStore from "./context/GlobalStore";

const { useContext } = React;

interface ContainerProps {
  children: React.ReactNode;
}
const Container = (props: ContainerProps): JSX.Element => {
  const { dispatch } = useContext(GlobalStore);

  React.useEffect(() => {
    request({
      method: "POST",
      action: "/api/login"
    }).then((res: any) => {
      if (res.RetCode === 0 && res.Data) {
        const { name, avatar_url } = res.Data[0];

        dispatch({
          type: "SIGN_IN",
          payload: {
            userName: name,
            avatar: avatar_url
          }
        })
      }
    })
  }, []);

  return (
    <div className="container">
      <Header/>
      <Background/>
      {props.children}
    </div>
  )
};

export default Container;
