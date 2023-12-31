=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.51.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.2.0-SNAPSHOT

=end

require 'date'
require 'time'

module LagoAPI
  class Timezone
    UTC = "UTC".freeze
    AFRICA_ALGIERS = "Africa/Algiers".freeze
    AFRICA_CAIRO = "Africa/Cairo".freeze
    AFRICA_CASABLANCA = "Africa/Casablanca".freeze
    AFRICA_HARARE = "Africa/Harare".freeze
    AFRICA_JOHANNESBURG = "Africa/Johannesburg".freeze
    AFRICA_MONROVIA = "Africa/Monrovia".freeze
    AFRICA_NAIROBI = "Africa/Nairobi".freeze
    AMERICA_ARGENTINA_BUENOS_AIRES = "America/Argentina/Buenos_Aires".freeze
    AMERICA_BOGOTA = "America/Bogota".freeze
    AMERICA_CARACAS = "America/Caracas".freeze
    AMERICA_CHICAGO = "America/Chicago".freeze
    AMERICA_CHIHUAHUA = "America/Chihuahua".freeze
    AMERICA_DENVER = "America/Denver".freeze
    AMERICA_GODTHAB = "America/Godthab".freeze
    AMERICA_GUATEMALA = "America/Guatemala".freeze
    AMERICA_GUYANA = "America/Guyana".freeze
    AMERICA_HALIFAX = "America/Halifax".freeze
    AMERICA_INDIANA_INDIANAPOLIS = "America/Indiana/Indianapolis".freeze
    AMERICA_JUNEAU = "America/Juneau".freeze
    AMERICA_LA_PAZ = "America/La_Paz".freeze
    AMERICA_LIMA = "America/Lima".freeze
    AMERICA_LOS_ANGELES = "America/Los_Angeles".freeze
    AMERICA_MAZATLAN = "America/Mazatlan".freeze
    AMERICA_MEXICO_CITY = "America/Mexico_City".freeze
    AMERICA_MONTERREY = "America/Monterrey".freeze
    AMERICA_MONTEVIDEO = "America/Montevideo".freeze
    AMERICA_NEW_YORK = "America/New_York".freeze
    AMERICA_PHOENIX = "America/Phoenix".freeze
    AMERICA_PUERTO_RICO = "America/Puerto_Rico".freeze
    AMERICA_REGINA = "America/Regina".freeze
    AMERICA_SANTIAGO = "America/Santiago".freeze
    AMERICA_SAO_PAULO = "America/Sao_Paulo".freeze
    AMERICA_ST_JOHNS = "America/St_Johns".freeze
    AMERICA_TIJUANA = "America/Tijuana".freeze
    ASIA_ALMATY = "Asia/Almaty".freeze
    ASIA_BAGHDAD = "Asia/Baghdad".freeze
    ASIA_BAKU = "Asia/Baku".freeze
    ASIA_BANGKOK = "Asia/Bangkok".freeze
    ASIA_CHONGQING = "Asia/Chongqing".freeze
    ASIA_COLOMBO = "Asia/Colombo".freeze
    ASIA_DHAKA = "Asia/Dhaka".freeze
    ASIA_HONG_KONG = "Asia/Hong_Kong".freeze
    ASIA_IRKUTSK = "Asia/Irkutsk".freeze
    ASIA_JAKARTA = "Asia/Jakarta".freeze
    ASIA_JERUSALEM = "Asia/Jerusalem".freeze
    ASIA_KABUL = "Asia/Kabul".freeze
    ASIA_KAMCHATKA = "Asia/Kamchatka".freeze
    ASIA_KARACHI = "Asia/Karachi".freeze
    ASIA_KATHMANDU = "Asia/Kathmandu".freeze
    ASIA_KOLKATA = "Asia/Kolkata".freeze
    ASIA_KRASNOYARSK = "Asia/Krasnoyarsk".freeze
    ASIA_KUALA_LUMPUR = "Asia/Kuala_Lumpur".freeze
    ASIA_KUWAIT = "Asia/Kuwait".freeze
    ASIA_MAGADAN = "Asia/Magadan".freeze
    ASIA_MUSCAT = "Asia/Muscat".freeze
    ASIA_NOVOSIBIRSK = "Asia/Novosibirsk".freeze
    ASIA_RANGOON = "Asia/Rangoon".freeze
    ASIA_RIYADH = "Asia/Riyadh".freeze
    ASIA_SEOUL = "Asia/Seoul".freeze
    ASIA_SHANGHAI = "Asia/Shanghai".freeze
    ASIA_SINGAPORE = "Asia/Singapore".freeze
    ASIA_SREDNEKOLYMSK = "Asia/Srednekolymsk".freeze
    ASIA_TAIPEI = "Asia/Taipei".freeze
    ASIA_TASHKENT = "Asia/Tashkent".freeze
    ASIA_TBILISI = "Asia/Tbilisi".freeze
    ASIA_TEHRAN = "Asia/Tehran".freeze
    ASIA_TOKYO = "Asia/Tokyo".freeze
    ASIA_ULAANBAATAR = "Asia/Ulaanbaatar".freeze
    ASIA_URUMQI = "Asia/Urumqi".freeze
    ASIA_VLADIVOSTOK = "Asia/Vladivostok".freeze
    ASIA_YAKUTSK = "Asia/Yakutsk".freeze
    ASIA_YEKATERINBURG = "Asia/Yekaterinburg".freeze
    ASIA_YEREVAN = "Asia/Yerevan".freeze
    ATLANTIC_AZORES = "Atlantic/Azores".freeze
    ATLANTIC_CAPE_VERDE = "Atlantic/Cape_Verde".freeze
    ATLANTIC_SOUTH_GEORGIA = "Atlantic/South_Georgia".freeze
    AUSTRALIA_ADELAIDE = "Australia/Adelaide".freeze
    AUSTRALIA_BRISBANE = "Australia/Brisbane".freeze
    AUSTRALIA_DARWIN = "Australia/Darwin".freeze
    AUSTRALIA_HOBART = "Australia/Hobart".freeze
    AUSTRALIA_MELBOURNE = "Australia/Melbourne".freeze
    AUSTRALIA_PERTH = "Australia/Perth".freeze
    AUSTRALIA_SYDNEY = "Australia/Sydney".freeze
    EUROPE_AMSTERDAM = "Europe/Amsterdam".freeze
    EUROPE_ATHENS = "Europe/Athens".freeze
    EUROPE_BELGRADE = "Europe/Belgrade".freeze
    EUROPE_BERLIN = "Europe/Berlin".freeze
    EUROPE_BRATISLAVA = "Europe/Bratislava".freeze
    EUROPE_BRUSSELS = "Europe/Brussels".freeze
    EUROPE_BUCHAREST = "Europe/Bucharest".freeze
    EUROPE_BUDAPEST = "Europe/Budapest".freeze
    EUROPE_COPENHAGEN = "Europe/Copenhagen".freeze
    EUROPE_DUBLIN = "Europe/Dublin".freeze
    EUROPE_HELSINKI = "Europe/Helsinki".freeze
    EUROPE_ISTANBUL = "Europe/Istanbul".freeze
    EUROPE_KALININGRAD = "Europe/Kaliningrad".freeze
    EUROPE_KIEV = "Europe/Kiev".freeze
    EUROPE_LISBON = "Europe/Lisbon".freeze
    EUROPE_LJUBLJANA = "Europe/Ljubljana".freeze
    EUROPE_LONDON = "Europe/London".freeze
    EUROPE_MADRID = "Europe/Madrid".freeze
    EUROPE_MINSK = "Europe/Minsk".freeze
    EUROPE_MOSCOW = "Europe/Moscow".freeze
    EUROPE_PARIS = "Europe/Paris".freeze
    EUROPE_PRAGUE = "Europe/Prague".freeze
    EUROPE_RIGA = "Europe/Riga".freeze
    EUROPE_ROME = "Europe/Rome".freeze
    EUROPE_SAMARA = "Europe/Samara".freeze
    EUROPE_SARAJEVO = "Europe/Sarajevo".freeze
    EUROPE_SKOPJE = "Europe/Skopje".freeze
    EUROPE_SOFIA = "Europe/Sofia".freeze
    EUROPE_STOCKHOLM = "Europe/Stockholm".freeze
    EUROPE_TALLINN = "Europe/Tallinn".freeze
    EUROPE_VIENNA = "Europe/Vienna".freeze
    EUROPE_VILNIUS = "Europe/Vilnius".freeze
    EUROPE_VOLGOGRAD = "Europe/Volgograd".freeze
    EUROPE_WARSAW = "Europe/Warsaw".freeze
    EUROPE_ZAGREB = "Europe/Zagreb".freeze
    EUROPE_ZURICH = "Europe/Zurich".freeze
    GMT12 = "GMT+12".freeze
    PACIFIC_APIA = "Pacific/Apia".freeze
    PACIFIC_AUCKLAND = "Pacific/Auckland".freeze
    PACIFIC_CHATHAM = "Pacific/Chatham".freeze
    PACIFIC_FAKAOFO = "Pacific/Fakaofo".freeze
    PACIFIC_FIJI = "Pacific/Fiji".freeze
    PACIFIC_GUADALCANAL = "Pacific/Guadalcanal".freeze
    PACIFIC_GUAM = "Pacific/Guam".freeze
    PACIFIC_HONOLULU = "Pacific/Honolulu".freeze
    PACIFIC_MAJURO = "Pacific/Majuro".freeze
    PACIFIC_MIDWAY = "Pacific/Midway".freeze
    PACIFIC_NOUMEA = "Pacific/Noumea".freeze
    PACIFIC_PAGO_PAGO = "Pacific/Pago_Pago".freeze
    PACIFIC_PORT_MORESBY = "Pacific/Port_Moresby".freeze
    PACIFIC_TONGATAPU = "Pacific/Tongatapu".freeze

    def self.all_vars
      @all_vars ||= [UTC, AFRICA_ALGIERS, AFRICA_CAIRO, AFRICA_CASABLANCA, AFRICA_HARARE, AFRICA_JOHANNESBURG, AFRICA_MONROVIA, AFRICA_NAIROBI, AMERICA_ARGENTINA_BUENOS_AIRES, AMERICA_BOGOTA, AMERICA_CARACAS, AMERICA_CHICAGO, AMERICA_CHIHUAHUA, AMERICA_DENVER, AMERICA_GODTHAB, AMERICA_GUATEMALA, AMERICA_GUYANA, AMERICA_HALIFAX, AMERICA_INDIANA_INDIANAPOLIS, AMERICA_JUNEAU, AMERICA_LA_PAZ, AMERICA_LIMA, AMERICA_LOS_ANGELES, AMERICA_MAZATLAN, AMERICA_MEXICO_CITY, AMERICA_MONTERREY, AMERICA_MONTEVIDEO, AMERICA_NEW_YORK, AMERICA_PHOENIX, AMERICA_PUERTO_RICO, AMERICA_REGINA, AMERICA_SANTIAGO, AMERICA_SAO_PAULO, AMERICA_ST_JOHNS, AMERICA_TIJUANA, ASIA_ALMATY, ASIA_BAGHDAD, ASIA_BAKU, ASIA_BANGKOK, ASIA_CHONGQING, ASIA_COLOMBO, ASIA_DHAKA, ASIA_HONG_KONG, ASIA_IRKUTSK, ASIA_JAKARTA, ASIA_JERUSALEM, ASIA_KABUL, ASIA_KAMCHATKA, ASIA_KARACHI, ASIA_KATHMANDU, ASIA_KOLKATA, ASIA_KRASNOYARSK, ASIA_KUALA_LUMPUR, ASIA_KUWAIT, ASIA_MAGADAN, ASIA_MUSCAT, ASIA_NOVOSIBIRSK, ASIA_RANGOON, ASIA_RIYADH, ASIA_SEOUL, ASIA_SHANGHAI, ASIA_SINGAPORE, ASIA_SREDNEKOLYMSK, ASIA_TAIPEI, ASIA_TASHKENT, ASIA_TBILISI, ASIA_TEHRAN, ASIA_TOKYO, ASIA_ULAANBAATAR, ASIA_URUMQI, ASIA_VLADIVOSTOK, ASIA_YAKUTSK, ASIA_YEKATERINBURG, ASIA_YEREVAN, ATLANTIC_AZORES, ATLANTIC_CAPE_VERDE, ATLANTIC_SOUTH_GEORGIA, AUSTRALIA_ADELAIDE, AUSTRALIA_BRISBANE, AUSTRALIA_DARWIN, AUSTRALIA_HOBART, AUSTRALIA_MELBOURNE, AUSTRALIA_PERTH, AUSTRALIA_SYDNEY, EUROPE_AMSTERDAM, EUROPE_ATHENS, EUROPE_BELGRADE, EUROPE_BERLIN, EUROPE_BRATISLAVA, EUROPE_BRUSSELS, EUROPE_BUCHAREST, EUROPE_BUDAPEST, EUROPE_COPENHAGEN, EUROPE_DUBLIN, EUROPE_HELSINKI, EUROPE_ISTANBUL, EUROPE_KALININGRAD, EUROPE_KIEV, EUROPE_LISBON, EUROPE_LJUBLJANA, EUROPE_LONDON, EUROPE_MADRID, EUROPE_MINSK, EUROPE_MOSCOW, EUROPE_PARIS, EUROPE_PRAGUE, EUROPE_RIGA, EUROPE_ROME, EUROPE_SAMARA, EUROPE_SARAJEVO, EUROPE_SKOPJE, EUROPE_SOFIA, EUROPE_STOCKHOLM, EUROPE_TALLINN, EUROPE_VIENNA, EUROPE_VILNIUS, EUROPE_VOLGOGRAD, EUROPE_WARSAW, EUROPE_ZAGREB, EUROPE_ZURICH, GMT12, PACIFIC_APIA, PACIFIC_AUCKLAND, PACIFIC_CHATHAM, PACIFIC_FAKAOFO, PACIFIC_FIJI, PACIFIC_GUADALCANAL, PACIFIC_GUAM, PACIFIC_HONOLULU, PACIFIC_MAJURO, PACIFIC_MIDWAY, PACIFIC_NOUMEA, PACIFIC_PAGO_PAGO, PACIFIC_PORT_MORESBY, PACIFIC_TONGATAPU].freeze
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      return value if Timezone.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #Timezone"
    end
  end
end
