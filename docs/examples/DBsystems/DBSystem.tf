resource "baremetal_database_db_system" "TFDBNode" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}" 
  compartment_id = "${var.compartment_ocid}"
  cpu_core_count = "${var.CPUCoreCount}"
  database_edition = "${var.DBEdition}"
  db_home {
    database {
      "admin_password" = "${var.DBAdminPassword}"
      "db_name" = "${var.DBName}"
    }
    db_version = "${var.DBVersion}"
    display_name = "${var.DBDisplayName}"
  }
  disk_redundancy = "${var.DBDiskRedundancy}"
  shape = "${var.DBNodeShape}"
  subnet_id = "${var.SubnetOCID}"
  ssh_public_keys = ["${var.ssh_private_key}"]
  display_name = "${var.DBNodeDisplayName}"
  domain = "${var.DBNodeDomainName}"
  hostname = "${var.DBNodeHostName}"
}
