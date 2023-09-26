-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 03 Agu 2021 pada 12.09
-- Versi server: 10.4.17-MariaDB
-- Versi PHP: 7.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_emi`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_invoices`
--

CREATE TABLE `tb_invoices` (
  `INVOICE_ID` int(11) NOT NULL,
  `PROJECT_ID` int(11) DEFAULT NULL,
  `CURRENCY` varchar(100) DEFAULT NULL,
  `TOTAL_AMOUNT` float DEFAULT NULL,
  `DISCOUNT` float DEFAULT NULL,
  `NOTES` varchar(100) DEFAULT NULL,
  `INVOICE_DUE_DATE` date DEFAULT NULL,
  `INVOICE_CREATE_DATE` date DEFAULT NULL,
  `INVOICE_TITLE` varchar(100) DEFAULT NULL,
  `ARRAY_OF_INVOICE_ITEM` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tb_invoices`
--

INSERT INTO `tb_invoices` (`INVOICE_ID`, `PROJECT_ID`, `CURRENCY`, `TOTAL_AMOUNT`, `DISCOUNT`, `NOTES`, `INVOICE_DUE_DATE`, `INVOICE_CREATE_DATE`, `INVOICE_TITLE`, `ARRAY_OF_INVOICE_ITEM`) VALUES
(1, 1, 'Rupiah', 1000000, 0, '-', '2021-08-09', '2021-08-03', 'Pembayaran Pertama', 'ItemName:Tes Item 2,Qty:1,Amount:100000,ItemName:Tes Item 4,Qty:2,Amount:500000'),
(3, 1, 'Rp', 1000000, 40000, 'Tes Ediitt', '2021-07-01', '2021-07-02', 'Judul', 'ItemName:Tes Item 2,Qty:1,Amount:100000,ItemName:Tes Item 4,Qty:2,Amount:500000'),
(4, 2, 'Rp', 1500000, 0, 'Tes DOang 2', '2021-07-01', '2021-07-02', 'Judul Dua', 'ItemName:Tes Item 3,Qty:1,Amount:1100000,ItemName:Tes Item 4,Qty:40,Amount:1500000'),
(5, 3, 'Rp', 1230000, 0, 'Tes DOang 3', '2021-07-01', '2021-07-02', 'Judul Dua', 'ItemName:Tes Item 5,Qty:1,Amount:1100000,ItemName:Tes Item 2,Qty:4,Amount:1500000'),
(6, 3, 'Rp', 1230000, 0, 'INI NOTES', '2021-07-01', '2021-07-02', 'Judul Dua', 'ItemName:Tes Item 7,Qty:7,Amount:1100000,ItemName:Tes Item 8,Qty:8,Amount:1500000'),
(7, 3, 'Rp', 1230000, 0, 'INI NOTES', '2021-07-01', '2021-07-02', 'Judul Dua', 'ItemName:Tes Item 5,Qty:1,Amount:1100000,ItemName:Tes Item 2,Qty:4,Amount:1500000');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_projects`
--

CREATE TABLE `tb_projects` (
  `PROJECT_ID` int(11) NOT NULL,
  `PROJECT_NAME` varchar(100) DEFAULT NULL,
  `COMPANY_NAME` varchar(100) DEFAULT NULL,
  `START_DATE` date DEFAULT NULL,
  `END_DATE` date DEFAULT NULL,
  `ARRAY_OF_USERID` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tb_projects`
--

INSERT INTO `tb_projects` (`PROJECT_ID`, `PROJECT_NAME`, `COMPANY_NAME`, `START_DATE`, `END_DATE`, `ARRAY_OF_USERID`) VALUES
(1, 'Project 1', 'Company 1', '2021-07-02', '2021-08-02', '1,2,6'),
(2, 'Project 12 Edit', 'Company 1', '2021-07-02', '2021-08-02', '1,2,3,4');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_users`
--

CREATE TABLE `tb_users` (
  `USER_ID` int(11) NOT NULL,
  `USERNAME` varchar(100) DEFAULT NULL,
  `EMAIL` varchar(100) DEFAULT NULL,
  `PASSWORD` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tb_users`
--

INSERT INTO `tb_users` (`USER_ID`, `USERNAME`, `EMAIL`, `PASSWORD`) VALUES
(1, 'SATU', 'SATU@GMAIL.COM', '123'),
(2, 'DUA', 'DUA@GMAIL.COM', 'DUA'),
(3, 'TIGA', 'TIGA@GMAIL.COM', 'TIGA');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `tb_invoices`
--
ALTER TABLE `tb_invoices`
  ADD PRIMARY KEY (`INVOICE_ID`);

--
-- Indeks untuk tabel `tb_projects`
--
ALTER TABLE `tb_projects`
  ADD PRIMARY KEY (`PROJECT_ID`);

--
-- Indeks untuk tabel `tb_users`
--
ALTER TABLE `tb_users`
  ADD PRIMARY KEY (`USER_ID`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `tb_invoices`
--
ALTER TABLE `tb_invoices`
  MODIFY `INVOICE_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `tb_projects`
--
ALTER TABLE `tb_projects`
  MODIFY `PROJECT_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `tb_users`
--
ALTER TABLE `tb_users`
  MODIFY `USER_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
