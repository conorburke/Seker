const url = {};

if (process.env.NODE_ENV === 'production') {
  url.api = 'http://localhost:5000'; // can be different than Dev if needed
} else if (process.env.NODE_ENV === 'development') {
  url.api = 'http://localhost:5000';
}

export default url;