import React from 'react';
import {Image} from 'reactwebcommon';
import {Panel} from 'reactpages';
import {connect} from 'react-redux';

class HeaderUI extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    let props = this.props;
    let module = props.module;
    let hs = module.properties.header;
    let settings = module.settings;
    let infoBlock = settings.infoBlock ? settings.infoBlock: "userBlock";
    console.log("header properties", hs, props, settings)
    return (
      <props.uikit.Block className={hs.className?hs.className:'header'}>
        <props.uikit.Block className="logo">
          {hs.image?<props.uikit.Block className="image"><Image src={hs.image}/></props.uikit.Block>:null}
          {hs.title?<props.uikit.Block className="title">{hs.title}</props.uikit.Block>:null}
        </props.uikit.Block>
        <Panel id={infoBlock} type="block" className="infoBlock"/>
      </props.uikit.Block>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    loggedIn: state.Security.status == "LoggedIn"
  }
}

const Header = connect(mapStateToProps)(HeaderUI);

export default Header
