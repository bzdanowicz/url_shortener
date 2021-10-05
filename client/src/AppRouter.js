import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import App from './App';


async function redirectOriginalUrl(path) {
    const apiUrl = process.env.REACT_APP_API_URL
    try {
      const response = await fetch(apiUrl + path, {
        method: 'get',
        headers: {
          'Accept': 'application/json, text/plain, */*',
          'Content-Type': 'application/json',
        }
      })
      const json = await response.json()
      const originalUrl = json["original_url"] || "/"
      window.location.href = originalUrl;
    } catch(error) {
        window.location.href = "/"
    }
}

function AppRouter() {
    return (
        <Router>
            <Switch>
                <Route exact path="/" component={App} />
                <Route path='/*' component={(props) => {
                    redirectOriginalUrl(props.location.pathname)
                    return null;
                }}/>
            </Switch>
        </Router>
    );
}

export default AppRouter;