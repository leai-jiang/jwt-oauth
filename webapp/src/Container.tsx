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
      action: "/api/rest/user"
    }).then((res: any) => {
      if (res.RetCode === 0) {
        const { name, avatar_url } = res.Data;

        dispatch({
          type: "SIGN_IN",
          payload: {
            userName: name,
            avatar: avatar_url,
            isSign: true
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
