resource "ec_dns_group" "failover1" {
  account_number = "4FDBB"
    group_type="cname"
    group_product_type="failover"
    name="failover-101"
    a {
      weight=100
      record {
        ttl=300
        rdata="10.10.1.3"
      }
    }
    a {
      weight=0
      record {
        ttl=300
        rdata="10.10.1.4"
      }
    }
}

resource "ec_dns_group" "loadbalancing1" {
  account_number = "4FDBB"
  group_type="cname"
  group_product_type="loadbalancing"
  name="loadbal-200"
  a {
    weight=33
    health_check {
      check_interval=300
      check_type_id=1
      content_verification="10"
      email_notification_address="notice@glory1.com"
      failed_check_threshold=10
      http_method_id=1
      # ip_address="" # IP address only required when check_type_id 3,4
      ip_version=1
      # port_number=80 # Port only required when check_type_id 3,4
      reintegration_method_id=1
      status= 4
      status_name="Unknown"
      uri="www.yahoo.com"
      timeout=100
    }
    record {
      ttl=300
      rdata="10.10.3.5"
    }
  }
  a {
    weight=33
    record {
      ttl=300
      rdata="10.10.3.6"
    }
  }
    a {
    weight=33
    record {
      ttl=300
      rdata="10.10.3.7"
    }
  }
}

resource "ec_dns_group" "loadbalancing2" {
  account_number = "4FDBB"
  group_type="cname"
  group_product_type="loadbalancing"
  name="loadbal-300"
  a {
    weight=33
    health_check {
      check_interval=300
      check_type_id=3
      content_verification="10"
      email_notification_address="notice@glory1.com"
      failed_check_threshold=10
      # http_method_id=1 # Only required with check_type_id 1,2
      ip_address="85.23.100.11"
      # ip_version=1 # Only required with check_type_id 1,2
      port_number=445
      reintegration_method_id=1
      status= 4
      status_name="Unknown"
      # uri="www.yahoo.com" # Only required with check_type_id 1,2
      timeout=100
    }
    record {
      ttl=300
      rdata="10.10.3.5"
    }
  }
  a {
    weight=33
    record {
      ttl=300
      rdata="10.10.3.6"
    }
  }
    a {
    weight=33
    record {
      ttl=300
      rdata="10.10.3.7"
    }
  }
}