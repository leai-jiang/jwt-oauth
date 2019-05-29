import * as React from "react";
import { Icon } from "antd";

export const Header = (): JSX.Element => {
	function loginGithub() {
		window.location.href = "/api/login/github"
	}

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
