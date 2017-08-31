using ReactNative.Bridge;
using System;
using System.Collections.Generic;
using Windows.ApplicationModel.Core;
using Windows.UI.Core;

namespace Com.Reactlibrary.RNLaatoonative
{
    /// <summary>
    /// A module that allows JS to share data.
    /// </summary>
    class RNLaatoonativeModule : NativeModuleBase
    {
        /// <summary>
        /// Instantiates the <see cref="RNLaatoonativeModule"/>.
        /// </summary>
        internal RNLaatoonativeModule()
        {

        }

        /// <summary>
        /// The name of the native module.
        /// </summary>
        public override string Name
        {
            get
            {
                return "RNLaatoonative";
            }
        }
    }
}
