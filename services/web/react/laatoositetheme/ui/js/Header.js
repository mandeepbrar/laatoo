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
    console.log("header properties", hs, props, settings)
    return (
      <_uikit.Block className={hs.className?hs.className:'header'}>
        <_uikit.Block  className="logo">
          <_uikit.Action action={{url:"/"}}>
            {hs.image?<_uikit.Block className="image"><Image src={hs.image}/></_uikit.Block>:null}
            {hs.title?<_uikit.Block className="title">{hs.title}</_uikit.Block>:null}
          </_uikit.Action>
        </_uikit.Block>
        {settings.infoBlock? <Panel id={settings.infoBlock} type="block" className="infoBlock"/> :null}
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
