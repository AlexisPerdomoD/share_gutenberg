import { NavLink } from "react-router-dom"

const NavBar = () => {
  return (
    <nav>
        <NavLink to="/collections">
            <h3>Collections</h3>
        </NavLink>
        <NavLink to="/profile">
            <h3>Profile</h3>
        </NavLink>
        <NavLink to="/about">
            <h3>Know more about Gutenberg Project</h3>
        </NavLink>
    </nav>
  )
}

export default NavBar

/* 
  className={({ isActive, isPending }) =>
    isPending ? "pending" : isActive ? "active" : ""
  }
*/