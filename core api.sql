CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `teknisi_id` int,
  `kerusakan_id` int,
  `jenis_hp` varchar(100),
  `jenis_platform` varchar(100),
  `status_service` enum('antrian', 'dikerjakan', 'selesai'),
  `lama_pengerjaan` varchar(100),
  `version` int
  `version_teknisi` int
);
