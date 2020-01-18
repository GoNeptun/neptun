-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Dec 25, 2019 at 09:12 PM
-- Server version: 5.7.26
-- PHP Version: 7.3.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `neptun`
--

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE IF NOT EXISTS `roles` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT 'User ID',
  `admin` int(11) NOT NULL COMMENT 'True or False',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='Users roles';

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `uid`, `admin`) VALUES
(1, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `site_settings`
--

CREATE TABLE IF NOT EXISTS `site_settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `maintenance` int(11) NOT NULL DEFAULT '1',
  `mail_confirmation` int(11) NOT NULL DEFAULT '0',
  `registration` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `site_settings`
--

INSERT INTO `site_settings` (`id`, `maintenance`, `mail_confirmation`, `registration`) VALUES
(1, 0, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `timehash`
--

CREATE TABLE IF NOT EXISTS `timehash` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `hash` varchar(254) NOT NULL DEFAULT '' COMMENT 'Hash',
  `name` varchar(254) DEFAULT NULL,
  `mail` varchar(254) DEFAULT '' COMMENT 'User’s e-mail address.',
  `param` varchar(254) NOT NULL,
  `created` int(11) NOT NULL DEFAULT '0' COMMENT 'Timestamp for when hash was created.',
  `deadline` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Stores time hashes.';

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT 'Unique user name.',
  `pass` varchar(128) NOT NULL DEFAULT '' COMMENT 'User’s password (hashed).',
  `mail` varchar(254) DEFAULT '' COMMENT 'User’s e-mail address.',
  `mailsent` int(11) DEFAULT '0' COMMENT 'Email Confirmation sent timestamp',
  `mail_confirmed` int(11) NOT NULL DEFAULT '0',
  `created` int(11) NOT NULL DEFAULT '0' COMMENT 'Timestamp for when user was created.',
  `access` int(11) NOT NULL DEFAULT '0' COMMENT 'Timestamp for previous time user accessed the site.',
  `login` int(11) NOT NULL DEFAULT '0' COMMENT 'Timestamp for user’s last login.',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'Whether the user is active(1) or blocked(0).',
  `completed` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `access` (`access`),
  KEY `created` (`created`),
  KEY `mail` (`mail`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='Stores user data.';

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `pass`, `mail`, `mailsent`, `mail_confirmed`, `created`, `access`, `login`, `status`, `completed`) VALUES
(1, 'admin', '$2a$08$m/iVl3UjEWD94poGUtN3oulpWwtmj/ZsBoWVrzBvC2BTpjrQ8Pgo.', 'admin@example.com', 0, 1, 0, 0, 0, 1, 0),
(2, 'user', '$2a$08$e6e2TH0byVKZJB8NfjmoPufN/Q3NgQYkIIXSMoZp4DFsK1flwjw6a', 'user@example.com', 0, 1, 0, 0, 0, 1, 0);
