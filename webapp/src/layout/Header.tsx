import * as React from "react";
import { Icon, Button } from "antd";

export const Header = (): JSX.Element => {
	return (
		<div className="header">
			<div className="logo">Sparta</div>
			<Icon type="user"/>
		</div>
	)
};
