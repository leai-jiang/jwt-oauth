import * as React from "react";
import { Icon } from "antd";

function loginGithub() {
	window.location.href = "http://localhost:5005/api/login/github"
}

export const Header = (): JSX.Element => {
	return (
		<div className="header">
			<div className="logo">Sparta</div>
			<div className="oauth-login">
				<Icon type="github" onClick={loginGithub} />
				<Icon type="wechat"/>
			</div>
		</div>
	)
};
