import * as React from "react";
import Background from "./background/Background";
import { Header } from "./layout"

interface ContainerProps {
    children: React.ReactNode;
}
const Container = (props: ContainerProps): JSX.Element => {
    return (
        <div>
            <Header/>
            <Background/>
            {props.children}
        </div>
    )
};

export default Container;
