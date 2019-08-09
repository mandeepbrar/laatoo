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
    console.log("header props", props)
    let hs = module.properties.header;
    let settings = module.settings;
    let infoBlock = settings.infoBlock ? settings.infoBlock: "userBlock";
    console.log("header properties", hs, props, settings)
    return (
      <_uikit.Block className={hs.className?hs.className:'header'}>
        <_uikit.Block className="logo">
          {hs.image?<_uikit.Block className="image"><Image src={hs.image}/></_uikit.Block>:null}
          {hs.title?<_uikit.Block className="title">{hs.title}</_uikit.Block>:null}
        </_uikit.Block>
        <Panel id={infoBlock} type="block" className="infoBlock"/>
      </_uikit.Block>
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
