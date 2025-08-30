import dynamic from "next/dynamic";
import React, { PropsWithChildren } from "react";

const NoSSR = (props: PropsWithChildren<{}>) => (
  <React.Fragment>{props.children}</React.Fragment>
);
export default dynamic(() => Promise.resolve(NoSSR), {
  ssr: false,
});
