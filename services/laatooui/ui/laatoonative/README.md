
# react-native-laatoonative

## Getting started

`$ npm install react-native-laatoonative --save`

### Mostly automatic installation

`$ react-native link react-native-laatoonative`

### Manual installation


#### iOS

1. In XCode, in the project navigator, right click `Libraries` ➜ `Add Files to [your project's name]`
2. Go to `node_modules` ➜ `react-native-laatoonative` and add `RNLaatoonative.xcodeproj`
3. In XCode, in the project navigator, select your project. Add `libRNLaatoonative.a` to your project's `Build Phases` ➜ `Link Binary With Libraries`
4. Run your project (`Cmd+R`)<

#### Android

1. Open up `android/app/src/main/java/[...]/MainActivity.java`
  - Add `import com.reactlibrary.RNLaatoonativePackage;` to the imports at the top of the file
  - Add `new RNLaatoonativePackage()` to the list returned by the `getPackages()` method
2. Append the following lines to `android/settings.gradle`:
  	```
  	include ':react-native-laatoonative'
  	project(':react-native-laatoonative').projectDir = new File(rootProject.projectDir, 	'../node_modules/react-native-laatoonative/android')
  	```
3. Insert the following lines inside the dependencies block in `android/app/build.gradle`:
  	```
      compile project(':react-native-laatoonative')
  	```

#### Windows
[Read it! :D](https://github.com/ReactWindows/react-native)

1. In Visual Studio add the `RNLaatoonative.sln` in `node_modules/react-native-laatoonative/windows/RNLaatoonative.sln` folder to their solution, reference from their app.
2. Open up your `MainPage.cs` app
  - Add `using Com.Reactlibrary.RNLaatoonative;` to the usings at the top of the file
  - Add `new RNLaatoonativePackage()` to the `List<IReactPackage>` returned by the `Packages` method


## Usage
```javascript
import RNLaatoonative from 'react-native-laatoonative';

// TODO: What to do with the module?
RNLaatoonative;
```
  