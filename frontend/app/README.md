# Book a Room

## Vite Library references
- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Running the application
To start the developer environment you need to install the packages and then start the localhost.
1. Install packages: ```npm install```
2. Start localhost: ```npm run dev```


## Orientation in repository
- You can find all components in ```./app/src/components``` in individual folders, each containing a ```.jsx``` file and ```.css``` file with the same name as the parent folder.
- All pages are listed in ```./app/src/pages``` with a ```.jsx``` file for each page.
- All assets (images, icons etc) can be found in the ```./app/src/assets``` folder.

## Generic Changes
If you want to edit the layout or other functions that are common for **ALL** pages you can perform these changes in ```./app/src/layout.jsx```.

## Rouing
The routing uses ```react-router-dom``` and is used since it's included with the React package.
### Add a route
1. Create a page (```.jsx``` file) in ```./app/src/pages```and export a ```React.JSX.Element``` object from this file.
2. Import this object into the ```./app/src/App.jsx``` file.
3. In the ```App.jsx``` file you can already see that there are a few routes and to add another one you add another ```<Route />``` attribute with the desired path and element. The element is the same object that you imported in step 2.