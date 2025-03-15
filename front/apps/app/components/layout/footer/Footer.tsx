import React from 'react';

const Footer = () => {
  return (
    <footer className="fixed bottom-0 left-0 w-full flex flex-col items-center justify-center p-4 bg-blue-900 text-white">
      <div className="text-center">
        <p className="text-sm">© 2025 Company Name. All rights reserved.</p>
      </div>
      <nav className="mt-2">
        <ul className="flex space-x-4">
          <li>
            <a href="#privacy" className="hover:text-gray-300">
              Privacy Policy
            </a>
          </li>
          <li>
            <a href="#terms" className="hover:text-gray-300">
              Terms of Service
            </a>
          </li>
          <li>
            <a href="#contact" className="hover:text-gray-300">
              Contact Us
            </a>
          </li>
        </ul>
      </nav>
      <div className="mt-2">
        <a href="https://facebook.com" className="hover:text-gray-300 mx-2">
          Facebook
        </a>
        <a href="https://twitter.com" className="hover:text-gray-300 mx-2">
          Twitter
        </a>
        <a href="https://instagram.com" className="hover:text-gray-300 mx-2">
          Instagram
        </a>
      </div>
    </footer>
  );
};

export default Footer;
