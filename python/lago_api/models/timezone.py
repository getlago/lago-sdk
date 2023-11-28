# coding: utf-8

"""
    Lago API documentation

    Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

    The version of the OpenAPI document: 0.52.0-beta
    Contact: tech@getlago.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
import pprint
import re  # noqa: F401
from enum import Enum



try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class Timezone(str, Enum):
    """
    Timezone
    """

    """
    allowed enum values
    """
    UTC = 'UTC'
    AFRICA_SLASH_ALGIERS = 'Africa/Algiers'
    AFRICA_SLASH_CAIRO = 'Africa/Cairo'
    AFRICA_SLASH_CASABLANCA = 'Africa/Casablanca'
    AFRICA_SLASH_HARARE = 'Africa/Harare'
    AFRICA_SLASH_JOHANNESBURG = 'Africa/Johannesburg'
    AFRICA_SLASH_MONROVIA = 'Africa/Monrovia'
    AFRICA_SLASH_NAIROBI = 'Africa/Nairobi'
    AMERICA_SLASH_ARGENTINA_SLASH_BUENOS_AIRES = 'America/Argentina/Buenos_Aires'
    AMERICA_SLASH_BOGOTA = 'America/Bogota'
    AMERICA_SLASH_CARACAS = 'America/Caracas'
    AMERICA_SLASH_CHICAGO = 'America/Chicago'
    AMERICA_SLASH_CHIHUAHUA = 'America/Chihuahua'
    AMERICA_SLASH_DENVER = 'America/Denver'
    AMERICA_SLASH_GODTHAB = 'America/Godthab'
    AMERICA_SLASH_GUATEMALA = 'America/Guatemala'
    AMERICA_SLASH_GUYANA = 'America/Guyana'
    AMERICA_SLASH_HALIFAX = 'America/Halifax'
    AMERICA_SLASH_INDIANA_SLASH_INDIANAPOLIS = 'America/Indiana/Indianapolis'
    AMERICA_SLASH_JUNEAU = 'America/Juneau'
    AMERICA_SLASH_LA_PAZ = 'America/La_Paz'
    AMERICA_SLASH_LIMA = 'America/Lima'
    AMERICA_SLASH_LOS_ANGELES = 'America/Los_Angeles'
    AMERICA_SLASH_MAZATLAN = 'America/Mazatlan'
    AMERICA_SLASH_MEXICO_CITY = 'America/Mexico_City'
    AMERICA_SLASH_MONTERREY = 'America/Monterrey'
    AMERICA_SLASH_MONTEVIDEO = 'America/Montevideo'
    AMERICA_SLASH_NEW_YORK = 'America/New_York'
    AMERICA_SLASH_PHOENIX = 'America/Phoenix'
    AMERICA_SLASH_PUERTO_RICO = 'America/Puerto_Rico'
    AMERICA_SLASH_REGINA = 'America/Regina'
    AMERICA_SLASH_SANTIAGO = 'America/Santiago'
    AMERICA_SLASH_SAO_PAULO = 'America/Sao_Paulo'
    AMERICA_SLASH_ST_JOHNS = 'America/St_Johns'
    AMERICA_SLASH_TIJUANA = 'America/Tijuana'
    ASIA_SLASH_ALMATY = 'Asia/Almaty'
    ASIA_SLASH_BAGHDAD = 'Asia/Baghdad'
    ASIA_SLASH_BAKU = 'Asia/Baku'
    ASIA_SLASH_BANGKOK = 'Asia/Bangkok'
    ASIA_SLASH_CHONGQING = 'Asia/Chongqing'
    ASIA_SLASH_COLOMBO = 'Asia/Colombo'
    ASIA_SLASH_DHAKA = 'Asia/Dhaka'
    ASIA_SLASH_HONG_KONG = 'Asia/Hong_Kong'
    ASIA_SLASH_IRKUTSK = 'Asia/Irkutsk'
    ASIA_SLASH_JAKARTA = 'Asia/Jakarta'
    ASIA_SLASH_JERUSALEM = 'Asia/Jerusalem'
    ASIA_SLASH_KABUL = 'Asia/Kabul'
    ASIA_SLASH_KAMCHATKA = 'Asia/Kamchatka'
    ASIA_SLASH_KARACHI = 'Asia/Karachi'
    ASIA_SLASH_KATHMANDU = 'Asia/Kathmandu'
    ASIA_SLASH_KOLKATA = 'Asia/Kolkata'
    ASIA_SLASH_KRASNOYARSK = 'Asia/Krasnoyarsk'
    ASIA_SLASH_KUALA_LUMPUR = 'Asia/Kuala_Lumpur'
    ASIA_SLASH_KUWAIT = 'Asia/Kuwait'
    ASIA_SLASH_MAGADAN = 'Asia/Magadan'
    ASIA_SLASH_MUSCAT = 'Asia/Muscat'
    ASIA_SLASH_NOVOSIBIRSK = 'Asia/Novosibirsk'
    ASIA_SLASH_RANGOON = 'Asia/Rangoon'
    ASIA_SLASH_RIYADH = 'Asia/Riyadh'
    ASIA_SLASH_SEOUL = 'Asia/Seoul'
    ASIA_SLASH_SHANGHAI = 'Asia/Shanghai'
    ASIA_SLASH_SINGAPORE = 'Asia/Singapore'
    ASIA_SLASH_SREDNEKOLYMSK = 'Asia/Srednekolymsk'
    ASIA_SLASH_TAIPEI = 'Asia/Taipei'
    ASIA_SLASH_TASHKENT = 'Asia/Tashkent'
    ASIA_SLASH_TBILISI = 'Asia/Tbilisi'
    ASIA_SLASH_TEHRAN = 'Asia/Tehran'
    ASIA_SLASH_TOKYO = 'Asia/Tokyo'
    ASIA_SLASH_ULAANBAATAR = 'Asia/Ulaanbaatar'
    ASIA_SLASH_URUMQI = 'Asia/Urumqi'
    ASIA_SLASH_VLADIVOSTOK = 'Asia/Vladivostok'
    ASIA_SLASH_YAKUTSK = 'Asia/Yakutsk'
    ASIA_SLASH_YEKATERINBURG = 'Asia/Yekaterinburg'
    ASIA_SLASH_YEREVAN = 'Asia/Yerevan'
    ATLANTIC_SLASH_AZORES = 'Atlantic/Azores'
    ATLANTIC_SLASH_CAPE_VERDE = 'Atlantic/Cape_Verde'
    ATLANTIC_SLASH_SOUTH_GEORGIA = 'Atlantic/South_Georgia'
    AUSTRALIA_SLASH_ADELAIDE = 'Australia/Adelaide'
    AUSTRALIA_SLASH_BRISBANE = 'Australia/Brisbane'
    AUSTRALIA_SLASH_DARWIN = 'Australia/Darwin'
    AUSTRALIA_SLASH_HOBART = 'Australia/Hobart'
    AUSTRALIA_SLASH_MELBOURNE = 'Australia/Melbourne'
    AUSTRALIA_SLASH_PERTH = 'Australia/Perth'
    AUSTRALIA_SLASH_SYDNEY = 'Australia/Sydney'
    EUROPE_SLASH_AMSTERDAM = 'Europe/Amsterdam'
    EUROPE_SLASH_ATHENS = 'Europe/Athens'
    EUROPE_SLASH_BELGRADE = 'Europe/Belgrade'
    EUROPE_SLASH_BERLIN = 'Europe/Berlin'
    EUROPE_SLASH_BRATISLAVA = 'Europe/Bratislava'
    EUROPE_SLASH_BRUSSELS = 'Europe/Brussels'
    EUROPE_SLASH_BUCHAREST = 'Europe/Bucharest'
    EUROPE_SLASH_BUDAPEST = 'Europe/Budapest'
    EUROPE_SLASH_COPENHAGEN = 'Europe/Copenhagen'
    EUROPE_SLASH_DUBLIN = 'Europe/Dublin'
    EUROPE_SLASH_HELSINKI = 'Europe/Helsinki'
    EUROPE_SLASH_ISTANBUL = 'Europe/Istanbul'
    EUROPE_SLASH_KALININGRAD = 'Europe/Kaliningrad'
    EUROPE_SLASH_KIEV = 'Europe/Kiev'
    EUROPE_SLASH_LISBON = 'Europe/Lisbon'
    EUROPE_SLASH_LJUBLJANA = 'Europe/Ljubljana'
    EUROPE_SLASH_LONDON = 'Europe/London'
    EUROPE_SLASH_MADRID = 'Europe/Madrid'
    EUROPE_SLASH_MINSK = 'Europe/Minsk'
    EUROPE_SLASH_MOSCOW = 'Europe/Moscow'
    EUROPE_SLASH_PARIS = 'Europe/Paris'
    EUROPE_SLASH_PRAGUE = 'Europe/Prague'
    EUROPE_SLASH_RIGA = 'Europe/Riga'
    EUROPE_SLASH_ROME = 'Europe/Rome'
    EUROPE_SLASH_SAMARA = 'Europe/Samara'
    EUROPE_SLASH_SARAJEVO = 'Europe/Sarajevo'
    EUROPE_SLASH_SKOPJE = 'Europe/Skopje'
    EUROPE_SLASH_SOFIA = 'Europe/Sofia'
    EUROPE_SLASH_STOCKHOLM = 'Europe/Stockholm'
    EUROPE_SLASH_TALLINN = 'Europe/Tallinn'
    EUROPE_SLASH_VIENNA = 'Europe/Vienna'
    EUROPE_SLASH_VILNIUS = 'Europe/Vilnius'
    EUROPE_SLASH_VOLGOGRAD = 'Europe/Volgograd'
    EUROPE_SLASH_WARSAW = 'Europe/Warsaw'
    EUROPE_SLASH_ZAGREB = 'Europe/Zagreb'
    EUROPE_SLASH_ZURICH = 'Europe/Zurich'
    GMT_PLUS_12 = 'GMT+12'
    PACIFIC_SLASH_APIA = 'Pacific/Apia'
    PACIFIC_SLASH_AUCKLAND = 'Pacific/Auckland'
    PACIFIC_SLASH_CHATHAM = 'Pacific/Chatham'
    PACIFIC_SLASH_FAKAOFO = 'Pacific/Fakaofo'
    PACIFIC_SLASH_FIJI = 'Pacific/Fiji'
    PACIFIC_SLASH_GUADALCANAL = 'Pacific/Guadalcanal'
    PACIFIC_SLASH_GUAM = 'Pacific/Guam'
    PACIFIC_SLASH_HONOLULU = 'Pacific/Honolulu'
    PACIFIC_SLASH_MAJURO = 'Pacific/Majuro'
    PACIFIC_SLASH_MIDWAY = 'Pacific/Midway'
    PACIFIC_SLASH_NOUMEA = 'Pacific/Noumea'
    PACIFIC_SLASH_PAGO_PAGO = 'Pacific/Pago_Pago'
    PACIFIC_SLASH_PORT_MORESBY = 'Pacific/Port_Moresby'
    PACIFIC_SLASH_TONGATAPU = 'Pacific/Tongatapu'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of Timezone from a JSON string"""
        return cls(json.loads(json_str))

