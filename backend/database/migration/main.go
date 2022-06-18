package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "kelasku.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`

CREATE TABLE IF NOT EXISTS siswa (
    id_siswa integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
	password varchar(255) not null,
	email varchar(255) not null,
	no_hp varchar(100) not null,
	nama varchar(255) not null,
	alamat text not null,
	tempat_lahir varchar(255) not null,
	tanggal_lahir date not null,
	jenis_kelamin varchar(255) not null,
	agama varchar(100) not null,
    file_name varchar(255) not null,
    show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS guru (
    id_guru integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
	password varchar(255) not null,
	nama varchar(255) not null,
	alamat text not null,
	no_hp varchar(255) not null,
	jenis_kelamin varchar(100) not null,
	agama varchar(100) not null,
	email varchar(255) not null,
	pendidikan varchar(255) not null,
	file_name varchar(255) not null,
    show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS wali_orangtua (
    id_wali integer not null primary key AUTOINCREMENT,
    id_siswa integer not null,
	username varchar(255) not null,
	password varchar(255) not null,
	nama varchar(255) not null,
	alamat text not null,
	no_hp varchar(255) not null,
	jenis_kelamin varchar(100) not null,
	email varchar(255) not null,
    show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS class (
    id_class integer not null primary key AUTOINCREMENT,
    id_guru integer not null,
	nama_kelas varchar(255) not null,
	keterangan text not null,
	kode_kelas varchar(255) not null,
	file_name varchar(255) not null,
	status_kelas TEXT CHECK( status_kelas IN ('0','1') ) not null DEFAULT '0',
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tugas (
    id_tugas integer not null primary key AUTOINCREMENT,
    id_class integer not null,
	judul_tugas varchar(255) not null,
	keterangan text not null,
	deadline datetime not null,
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS detail_tugas (
    id integer not null primary key AUTOINCREMENT,
    id_tugas integer not null,
	id_siswa integer not null,
	status_pengerjaan varchar(255) not null,
	waktu_pengumpulan datetime not null,
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS jadwal (
    id_jadwal integer not null primary key AUTOINCREMENT,
    id_class integer not null,
	hari varchar(255) not null,
	tanggal date not null,
	jam varchar(255) not null,
	status TEXT CHECK( status IN ('0','1','2') ) not null,
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS dt_kelassiswa (
    id_dt_kelassiswa integer not null primary key AUTOINCREMENT,
    id_siswa integer not null,
	id_class integer not null,
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS presensi (
    id_presensi integer not null primary key AUTOINCREMENT,
    id_dt_kelassiswa integer not null,
	tanggal date not null,
	jam varchar(255) not null,
	status_hadir varchar(255) not null,
	show_item TEXT CHECK( show_item IN ('0','1') )  not null DEFAULT '1',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`)

	if err != nil {
		panic(err)
	}
}
