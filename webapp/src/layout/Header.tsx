import * as React from "react";
import { Avatar } from "antd"
import GlobalStore from "../context/GlobalStore";
import signIn from "../components/SignDialog";
import request from "../utils/request";

const { useContext } = React;

export const Header = (): JSX.Element => {
  const { state, dispatch } = useContext(GlobalStore);

  function signOut() {
  	request({
			method: "POST",
			action: "/api/logout",
		}).then((res: any) => {
			dispatch({ type: "SIGN_OUT" });
		});
	}

	return (
		<div className="header">
			<div className="logo">Sparta</div>
	  {
		state.isSign ? (
		  <div>
			<Avatar src={state.avatar}/>
			<em style={{ color: "#FFF", margin: "0 10px" }}>{state.userName}</em>
						<a href="javascript:void(0)" onClick={() => signOut()}>Sign out</a>
		  </div>
		) : <a href="javascript:void(0)" onClick={() => signIn()}>Sign in</a>
	  }
		</div>
	)
};
